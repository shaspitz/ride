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

var _ types.QueryServer = Keeper{}

func (k Keeper) NextRide(c context.Context, req *types.QueryGetNextRideRequest) (*types.QueryGetNextRideResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNextRide(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNextRideResponse{NextRide: val}, nil
}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

func (k Keeper) StoredRideAll(c context.Context, req *types.QueryAllStoredRideRequest) (*types.QueryAllStoredRideResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var storedRides []types.StoredRide
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	storedRideStore := prefix.NewStore(store, types.KeyPrefix(types.StoredRideKeyPrefix))

	pageRes, err := query.Paginate(storedRideStore, req.Pagination, func(key []byte, value []byte) error {
		var storedRide types.StoredRide
		if err := k.cdc.Unmarshal(value, &storedRide); err != nil {
			return err
		}

		storedRides = append(storedRides, storedRide)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStoredRideResponse{StoredRide: storedRides, Pagination: pageRes}, nil
}

func (k Keeper) StoredRide(c context.Context, req *types.QueryGetStoredRideRequest) (*types.QueryGetStoredRideResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetStoredRide(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStoredRideResponse{StoredRide: val}, nil
}

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
