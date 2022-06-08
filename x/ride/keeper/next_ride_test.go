package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/smarshall-spitzbart/ride/testutil/keeper"
	"github.com/smarshall-spitzbart/ride/testutil/nullify"
	"github.com/smarshall-spitzbart/ride/x/ride/keeper"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
)

func createTestNextRide(keeper *keeper.Keeper, ctx sdk.Context) types.NextRide {
	item := types.NextRide{}
	keeper.SetNextRide(ctx, item)
	return item
}

func TestNextRideGet(t *testing.T) {
	keeper, ctx := keepertest.RideKeeper(t)
	item := createTestNextRide(keeper, ctx)
	rst, found := keeper.GetNextRide(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestNextRideRemove(t *testing.T) {
	keeper, ctx := keepertest.RideKeeper(t)
	createTestNextRide(keeper, ctx)
	keeper.RemoveNextRide(ctx)
	_, found := keeper.GetNextRide(ctx)
	require.False(t, found)
}
