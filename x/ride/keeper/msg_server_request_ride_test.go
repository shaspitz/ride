package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/smarshall-spitzbart/ride/testutil/keeper"
	"github.com/smarshall-spitzbart/ride/x/ride"
	"github.com/smarshall-spitzbart/ride/x/ride/keeper"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
	"github.com/stretchr/testify/require"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
)

func setupMsgServerRequestRide(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.RideKeeper(t)
	ride.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

func TestRequestRide(t *testing.T) {
	msgServer, _, context := setupMsgServerRequestRide(t)
	createResponse, err := msgServer.RequestRide(context, &types.MsgRequestRide{
		Creator:       alice, // Creator is passenger.
		StartLocation: "some lat/long",
		Destination:   "some other lat/long",
		MutualStake:   50,
		HourlyPay:     25,
		DistanceTip:   10,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgRequestRideResponse{
		IdValue: "1",
	}, *createResponse)
}

// Tests storage values when requesting a ride.
func TestRideRequestStorage(t *testing.T) {
	msgServer, keeper, context := setupMsgServerRequestRide(t)
	msgServer.RequestRide(context, &types.MsgRequestRide{
		Creator:       alice,
		StartLocation: "some loc",
		Destination:   "some dest",
		MutualStake:   30,
		HourlyPay:     15,
		DistanceTip:   5,
	})
	// Ensure IdValue counter is incremented.
	nextGame, found := keeper.GetNextRide(sdk.UnwrapSDKContext(context))
	require.True(t, found)
	require.EqualValues(t, types.NextRide{
		IdValue: 2,
	}, nextGame)
	// Ensure stored ride can be accessed from first value.
	game1, found1 := keeper.GetStoredRide(sdk.UnwrapSDKContext(context), "1")
	require.True(t, found1)
	require.EqualValues(t, types.StoredRide{
		Index:       "1",
		Destination: "some dest",
		Driver:      "", // Driver should not be set yet.
		Passenger:   alice,
		MutualStake: 30,
		PayPerHour:  15,
		DistanceTip: 5,
	}, game1)
}

// TODO: Make tests for the functionality that'll be included in "ValidateRequestRide"
