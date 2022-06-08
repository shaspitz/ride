package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/smarshall-spitzbart/ride/testutil/keeper"
	"github.com/smarshall-spitzbart/ride/testutil/nullify"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
)

func TestNextRideQuery(t *testing.T) {
	keeper, ctx := keepertest.RideKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestNextRide(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetNextRideRequest
		response *types.QueryGetNextRideResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetNextRideRequest{},
			response: &types.QueryGetNextRideResponse{NextRide: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.NextRide(wctx, tc.request)
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
