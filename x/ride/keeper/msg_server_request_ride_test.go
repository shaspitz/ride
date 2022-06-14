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
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func setupMsgServerRequestRide(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.RideKeeper(t)
	ride.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

func TestRequestRide(t *testing.T) {
	msgServer, _, context := setupMsgServerRequestRide(t)
	reqResponse, err := msgServer.RequestRide(context, &types.MsgRequestRide{
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
	}, *reqResponse)
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
	nextRide, found := keeper.GetNextRide(sdk.UnwrapSDKContext(context))
	require.True(t, found)
	require.EqualValues(t, types.NextRide{
		IdValue: 2,
	}, nextRide)
	// Ensure stored ride can be accessed from key.
	ride1, found1 := keeper.GetStoredRide(sdk.UnwrapSDKContext(context), "1")
	require.True(t, found1)
	require.EqualValues(t, types.StoredRide{
		Index:       "1",
		Destination: "some dest",
		Driver:      "", // Driver should not be set yet.
		Passenger:   alice,
		MutualStake: 30,
		PayPerHour:  15,
		DistanceTip: 5,
	}, ride1)
	require.Empty(t, ride1.AcceptanceTime)
	require.Empty(t, ride1.FinishLocation)
	require.Empty(t, ride1.FinishTime)
	require.False(t, ride1.HasAssignedDriver())
	require.False(t, ride1.IsFinished())
}

// TODO: Make tests for the functionality that'll be included in "ValidateRequestRide"

func TestRequestRideEventEmitted(t *testing.T) {
	msgServer, _, context := setupMsgServerRequestRide(t)
	msgServer.RequestRide(context, &types.MsgRequestRide{
		Creator:       alice,
		StartLocation: "some loc",
		Destination:   "some dest",
		MutualStake:   30,
		HourlyPay:     15,
		DistanceTip:   5,
	})
	ctx := sdk.UnwrapSDKContext(context)
	require.NotNil(t, ctx)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 1)
	event := events[0]
	require.EqualValues(t, sdk.StringEvent{
		Type: "message",
		Attributes: []sdk.Attribute{
			{Key: "module", Value: "ride"},
			{Key: "action", Value: "NewRideRequest"},
			{Key: "Passenger", Value: alice},
			{Key: "Index", Value: "1"},
			{Key: "StartLocation", Value: "some loc"},
		},
	}, event)
}
