package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/smarshall-spitzbart/ride/testutil/keeper"
	"github.com/smarshall-spitzbart/ride/testutil/nullify"
	"github.com/smarshall-spitzbart/ride/x/ride/keeper"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStoredRide(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StoredRide {
	items := make([]types.StoredRide, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetStoredRide(ctx, items[i])
	}
	return items
}

func TestStoredRideGet(t *testing.T) {
	keeper, ctx := keepertest.RideKeeper(t)
	items := createNStoredRide(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStoredRide(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStoredRideRemove(t *testing.T) {
	keeper, ctx := keepertest.RideKeeper(t)
	items := createNStoredRide(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStoredRide(ctx,
			item.Index,
		)
		_, found := keeper.GetStoredRide(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestStoredRideGetAll(t *testing.T) {
	keeper, ctx := keepertest.RideKeeper(t)
	items := createNStoredRide(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStoredRide(ctx)),
	)
}
