package ride_test

import (
	"testing"

	keepertest "github.com/smarshall-spitzbart/ride/testutil/keeper"
	"github.com/smarshall-spitzbart/ride/testutil/nullify"
	"github.com/smarshall-spitzbart/ride/x/ride"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		NextRide: &types.NextRide{
			IdValue: 69,
		},
		StoredRideList: []types.StoredRide{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RideKeeper(t)
	ride.InitGenesis(ctx, *k, genesisState)
	got := ride.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.NextRide, got.NextRide)
	require.ElementsMatch(t, genesisState.StoredRideList, got.StoredRideList)
	// this line is used by starport scaffolding # genesis/test/assert
}
