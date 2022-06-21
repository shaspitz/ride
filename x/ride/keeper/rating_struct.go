package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
)

// SetRatingStruct set a specific ratingStruct in the store from its index
func (k Keeper) SetRatingStruct(ctx sdk.Context, ratingStruct types.RatingStruct) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RatingStructKeyPrefix))
	b := k.cdc.MustMarshal(&ratingStruct)
	store.Set(types.RatingStructKey(
		ratingStruct.Index,
	), b)
}

// GetRatingStruct returns a ratingStruct from its index
func (k Keeper) GetRatingStruct(
	ctx sdk.Context,
	index string,

) (val types.RatingStruct, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RatingStructKeyPrefix))

	b := store.Get(types.RatingStructKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRatingStruct removes a ratingStruct from the store
func (k Keeper) RemoveRatingStruct(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RatingStructKeyPrefix))
	store.Delete(types.RatingStructKey(
		index,
	))
}

// GetAllRatingStruct returns all ratingStruct
func (k Keeper) GetAllRatingStruct(ctx sdk.Context) (list []types.RatingStruct) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RatingStructKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RatingStruct
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
