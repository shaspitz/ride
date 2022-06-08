package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
)

// SetNextRide set nextRide in the store
func (k Keeper) SetNextRide(ctx sdk.Context, nextRide types.NextRide) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextRideKey))
	b := k.cdc.MustMarshal(&nextRide)
	store.Set([]byte{0}, b)
}

// GetNextRide returns nextRide
func (k Keeper) GetNextRide(ctx sdk.Context) (val types.NextRide, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextRideKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNextRide removes nextRide from the store
func (k Keeper) RemoveNextRide(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextRideKey))
	store.Delete([]byte{0})
}
