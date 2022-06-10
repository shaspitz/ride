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
