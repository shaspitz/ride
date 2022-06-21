package ride

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/smarshall-spitzbart/ride/x/ride/keeper"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.NextRide != nil {
		k.SetNextRide(ctx, *genState.NextRide)
	}
	// Set all the storedRide
	for _, elem := range genState.StoredRideList {
		k.SetStoredRide(ctx, elem)
	}
	// Set all the ratingStruct
	for _, elem := range genState.RatingStructList {
		k.SetRatingStruct(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all nextRide
	nextRide, found := k.GetNextRide(ctx)
	if found {
		genesis.NextRide = &nextRide
	}
	genesis.StoredRideList = k.GetAllStoredRide(ctx)
	genesis.RatingStructList = k.GetAllRatingStruct(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
