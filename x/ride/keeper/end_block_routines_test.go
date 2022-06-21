package keeper_test

import (
	"context"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/smarshall-spitzbart/ride/x/ride/keeper"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
	"github.com/stretchr/testify/require"
)

func setupExpiredRideRequest(t testing.TB, msgServer types.MsgServer,
	k keeper.Keeper, context context.Context) {
	reqResponse, err := msgServer.RequestRide(context, &types.MsgRequestRide{
		Creator:       alice,
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
	ctx := sdk.UnwrapSDKContext(context)
	ride1, found1 := k.GetStoredRide(ctx, "1")
	require.True(t, found1)
	ride1.Deadline = types.TimeToString(ctx.BlockTime().Add(time.Duration(-1)))
	k.SetStoredRide(ctx, ride1)
}

func setupExpiredActiveRide(t testing.TB, msgServer types.MsgServer,
	k keeper.Keeper, context context.Context) {

	setupExpiredRideRequest(t, msgServer, k, context)
	ctx := sdk.UnwrapSDKContext(context)

	acceptRideResponse, err := msgServer.Accept(context, &types.MsgAccept{
		Creator: bob,
		IdValue: "1",
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgAcceptResponse{
		Success: true,
	}, *acceptRideResponse)

	// Re-expire ride, since acceptance handler will signify activity.
	ride1, found1 := k.GetStoredRide(sdk.UnwrapSDKContext(context), "1")
	ride1.Deadline = types.TimeToString(ctx.BlockTime().Add(time.Duration(-1)))
	k.SetStoredRide(ctx, ride1)

	require.True(t, found1)
	require.True(t, ride1.HasAssignedDriver())
	exp, err := ride1.HasExpired(sdk.UnwrapSDKContext(context))
	require.Nil(t, err)
	require.True(t, exp)
}

func setupExpiredFinishedRide(t testing.TB, msgServer types.MsgServer,
	k keeper.Keeper, context context.Context) {

	setupExpiredActiveRide(t, msgServer, k, context)
	ctx := sdk.UnwrapSDKContext(context)

	finishRideResponse, err := msgServer.Finish(context, &types.MsgFinish{
		Creator:  bob,
		IdValue:  "1",
		Location: "some finish loc",
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgFinishResponse{
		Success: true,
	}, *finishRideResponse)

	// Re-expire ride, since finish handler will signify activity.
	ride1, found1 := k.GetStoredRide(sdk.UnwrapSDKContext(context), "1")
	ride1.Deadline = types.TimeToString(ctx.BlockTime().Add(time.Duration(-1)))
	k.SetStoredRide(ctx, ride1)

	require.True(t, found1)
	require.True(t, ride1.IsFinished())
	exp, err := ride1.HasExpired(sdk.UnwrapSDKContext(context))
	require.Nil(t, err)
	require.True(t, exp)
}

func TestExpiredRideRequest(t *testing.T) {
	msgServer, k, context := setupMsgServerWithDefaultGenesis(t)
	ctx := sdk.UnwrapSDKContext(context)
	setupExpiredRideRequest(t, msgServer, k, context)

	k.CleanupExpiredRides(context)

	// Cleanup handler should have removed ride from store, and removed it from FIFO.
	_, found1 := k.GetStoredRide(ctx, "1")
	require.False(t, found1)
	nextRide, found2 := k.GetNextRide(ctx)
	require.True(t, found2)
	require.EqualValues(t, types.NoFifoIdKey, nextRide.FifoHead)
	require.EqualValues(t, types.NoFifoIdKey, nextRide.FifoTail)

	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 1)
	event := events[0]
	require.EqualValues(t, []sdk.Attribute{
		{Key: "module", Value: "ride"},
		{Key: "action", Value: "RideRequestExpired"},
		// Throw out attributes 0 through 4, coming from ride request event itself.
		// We're only testing the expiration event here.
	}, event.Attributes[5:])

	// TODO: Test funds being returned to passenger here.
}

func TestExpiredActiveRide(t *testing.T) {
	msgServer, k, context := setupMsgServerWithDefaultGenesis(t)
	ctx := sdk.UnwrapSDKContext(context)
	setupExpiredActiveRide(t, msgServer, k, context)

	k.CleanupExpiredRides(context)

	// Cleanup handler should have transformed the stored ride into a finished ride
	// with it's deadline reset.
	storedRide, found := k.GetStoredRide(ctx, "1")
	require.True(t, found)
	require.True(t, storedRide.IsFinished())
	require.EqualValues(t, types.TimeToString(ctx.BlockTime().Add(2*time.Minute)), storedRide.Deadline)
	require.EqualValues(t, types.TimeToString(ctx.BlockTime()), storedRide.FinishTime)
	require.EqualValues(t, "unknown", storedRide.FinishLocation)

	// Stored ride is still present in the FIFO linked list.
	nextRide, found2 := k.GetNextRide(ctx)
	require.True(t, found2)
	require.EqualValues(t, "1", nextRide.FifoHead)
	require.EqualValues(t, "1", nextRide.FifoTail)

	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 1)
	event := events[0]
	require.EqualValues(t, []sdk.Attribute{
		{Key: "module", Value: "ride"},
		{Key: "action", Value: "ActiveRideExpired"},
		// Throw out attributes 0 through 8, coming from ride request and acceptance events.
		// We're only testing the expiration event here.
	}, event.Attributes[9:])
}

func TestExpiredFinishedRide(t *testing.T) {
	msgServer, k, context := setupMsgServerWithDefaultGenesis(t)
	ctx := sdk.UnwrapSDKContext(context)
	setupExpiredFinishedRide(t, msgServer, k, context)

	k.CleanupExpiredRides(context)

	// Cleanup handler should have removed ride from store, and removed it from FIFO.
	_, found1 := k.GetStoredRide(ctx, "1")
	require.False(t, found1)
	nextRide, found2 := k.GetNextRide(ctx)
	require.True(t, found2)
	require.EqualValues(t, types.NoFifoIdKey, nextRide.FifoHead)
	require.EqualValues(t, types.NoFifoIdKey, nextRide.FifoTail)

	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 1)
	event := events[0]
	require.EqualValues(t, []sdk.Attribute{
		{Key: "module", Value: "ride"},
		{Key: "action", Value: "FinishedRideExpired"},
		// Throw out attributes 0 through 10, coming from ride request, acceptance, and finish events.
		// We're only testing the expiration event here.
	}, event.Attributes[11:])
}
