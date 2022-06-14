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

func setupMsgServerAcceptRide(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.RideKeeper(t)
	ride.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)
	// Request ride within setup function for ease.
	server.RequestRide(context, &types.MsgRequestRide{
		Creator:       alice,
		StartLocation: "some loc",
		Destination:   "some other loc",
		MutualStake:   50,
		HourlyPay:     15,
		DistanceTip:   10,
	})
	return server, *k, context
}

func TestAcceptRide(t *testing.T) {
	msgServer, _, context := setupMsgServerAcceptRide(t)
	acceptRideResponse, err := msgServer.Accept(context, &types.MsgAccept{
		Creator: bob,
		IdValue: "1",
	})
	// Require success in msg server handling.
	require.Nil(t, err)
	require.EqualValues(t, types.MsgAcceptResponse{
		Success: true,
	}, *acceptRideResponse)
}

func TestAcceptRideStorage(t *testing.T) {
	msgServer, keeper, context := setupMsgServerAcceptRide(t)
	acceptRideResponse, err := msgServer.Accept(context, &types.MsgAccept{
		Creator: bob,
		IdValue: "1",
	})
	// Require success in msg server handling.
	require.Nil(t, err)
	require.EqualValues(t, types.MsgAcceptResponse{
		Success: true,
	}, *acceptRideResponse)
	// Ensure stored ride can be accessed from key.
	ride1, found1 := keeper.GetStoredRide(sdk.UnwrapSDKContext(context), "1")
	require.True(t, found1)
	// Ensure no other fields were mutated, and that driver is now assigned.
	require.EqualValues(t, types.StoredRide{
		Index:       "1",
		Destination: "some other loc",
		Driver:      bob,
		Passenger:   alice,
		MutualStake: 50,
		PayPerHour:  15,
		DistanceTip: 10,
		// Bypass check for acceptance time, test that it's populated below.
		AcceptanceTime: ride1.AcceptanceTime,
	}, ride1)
	require.NotEmpty(t, ride1.AcceptanceTime)
	require.Empty(t, ride1.FinishTime)
	require.Empty(t, ride1.FinishLocation)
	require.True(t, ride1.HasAssignedDriver())
	require.False(t, ride1.IsFinished())
	// Finally, ensure that another driver cannot accept the ride.
	acceptRideResponse, err = msgServer.Accept(context, &types.MsgAccept{
		Creator: carol,
		IdValue: "1",
	})
	require.NotNil(t, err)
	require.EqualValues(t, types.MsgAcceptResponse{
		Success: false,
	}, *acceptRideResponse)
}

// TODO: Make tests for the functionality that'll be included in "ValidateAcceptRide"

func TestAcceptRideEventEmitted(t *testing.T) {
	msgServer, _, context := setupMsgServerAcceptRide(t)
	msgServer.Accept(context, &types.MsgAccept{
		Creator: bob,
		IdValue: "1",
	})
	ctx := sdk.UnwrapSDKContext(context)
	require.NotNil(t, ctx)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 1)
	event := events[0]
	require.EqualValues(t, []sdk.Attribute{
		{Key: "module", Value: "ride"},
		{Key: "action", Value: "RideAccepted"},
		{Key: "Driver", Value: bob},
		{Key: "IdValue", Value: "1"},
		// Throw out attributes 0 through 4, coming from request ride event.
	}, event.Attributes[5:])
}
