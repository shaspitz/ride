package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/smarshall-spitzbart/ride/testutil/keeper"
	"github.com/smarshall-spitzbart/ride/testutil/nullify"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestRatingStructQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.RideKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRatingStruct(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRatingStructRequest
		response *types.QueryGetRatingStructResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetRatingStructRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetRatingStructResponse{RatingStruct: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetRatingStructRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetRatingStructResponse{RatingStruct: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetRatingStructRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.RatingStruct(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestRatingStructQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.RideKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRatingStruct(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllRatingStructRequest {
		return &types.QueryAllRatingStructRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RatingStructAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RatingStruct), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RatingStruct),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RatingStructAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RatingStruct), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RatingStruct),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.RatingStructAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.RatingStruct),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.RatingStructAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
