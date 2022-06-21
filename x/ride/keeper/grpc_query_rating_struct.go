package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RatingStructAll(c context.Context, req *types.QueryAllRatingStructRequest) (*types.QueryAllRatingStructResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var ratingStructs []types.RatingStruct
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	ratingStructStore := prefix.NewStore(store, types.KeyPrefix(types.RatingStructKeyPrefix))

	pageRes, err := query.Paginate(ratingStructStore, req.Pagination, func(key []byte, value []byte) error {
		var ratingStruct types.RatingStruct
		if err := k.cdc.Unmarshal(value, &ratingStruct); err != nil {
			return err
		}

		ratingStructs = append(ratingStructs, ratingStruct)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRatingStructResponse{RatingStruct: ratingStructs, Pagination: pageRes}, nil
}

func (k Keeper) RatingStruct(c context.Context, req *types.QueryGetRatingStructRequest) (*types.QueryGetRatingStructResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRatingStruct(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRatingStructResponse{RatingStruct: val}, nil
}
