package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
)

// SetStoredRide set a specific storedRide in the store from its index
func (k Keeper) SetStoredRide(ctx sdk.Context, storedRide types.StoredRide) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredRideKeyPrefix))
	b := k.cdc.MustMarshal(&storedRide)
	store.Set(types.StoredRideKey(
		storedRide.Index,
	), b)
}

// GetStoredRide returns a storedRide from its index
func (k Keeper) GetStoredRide(
	ctx sdk.Context,
	index string,

) (val types.StoredRide, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredRideKeyPrefix))

	b := store.Get(types.StoredRideKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStoredRide removes a storedRide from the store
func (k Keeper) RemoveStoredRide(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredRideKeyPrefix))
	store.Delete(types.StoredRideKey(
		index,
	))
}

// GetAllStoredRide returns all storedRide
func (k Keeper) GetAllStoredRide(ctx sdk.Context) (list []types.StoredRide) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredRideKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StoredRide
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
