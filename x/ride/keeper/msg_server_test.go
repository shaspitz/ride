package keeper_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

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

func setupMsgServerWithDefaultGenesis(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.RideKeeper(t)
	ride.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

func TestRequestRide(t *testing.T) {
	msgServer, _, context := setupMsgServerWithDefaultGenesis(t)
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
	msgServer, keeper, context := setupMsgServerWithDefaultGenesis(t)
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
		IdValue:  2,
		FifoHead: "1",
		FifoTail: "1",
	}, nextRide)
	// Ensure stored ride can be accessed from key.
	ride1, found1 := keeper.GetStoredRide(sdk.UnwrapSDKContext(context), "1")
	require.True(t, found1)
	require.EqualValues(t, types.StoredRide{
		Index:            "1",
		Destination:      "some dest",
		DriverAddress:    "", // Driver should not be set yet.
		PassengerAddress: alice,
		MutualStake:      30,
		PayPerHour:       15,
		DistanceTip:      5,
		BeforeId:         "-1",
		AfterId:          "-1",
		Deadline:         types.TimeToString(sdk.UnwrapSDKContext(context).BlockTime().Add(types.DeadlinePeriod)),
	}, ride1)
	require.Empty(t, ride1.AcceptanceTime)
	require.Empty(t, ride1.FinishLocation)
	require.Empty(t, ride1.FinishTime)
	require.False(t, ride1.HasAssignedDriver())
	require.False(t, ride1.IsFinished())
}

func TestRideRequestValidation(t *testing.T) {
	msgServer, _, context := setupMsgServerWithDefaultGenesis(t)

	// mutual stake must be > distance tip + 1 hour pay
	_, err := msgServer.RequestRide(context, &types.MsgRequestRide{
		Creator:       alice,
		StartLocation: "a loc",
		Destination:   "a dest",
		HourlyPay:     5,
		DistanceTip:   3,
		MutualStake:   4,
	})
	require.NotNil(t, err)

	_, err = msgServer.RequestRide(context, &types.MsgRequestRide{
		Creator:       alice,
		StartLocation: "a loc",
		Destination:   "a dest",
		HourlyPay:     5,
		DistanceTip:   3,
		MutualStake:   2,
	})
	require.NotNil(t, err)

	_, err = msgServer.RequestRide(context, &types.MsgRequestRide{
		Creator:       alice,
		StartLocation: "a loc",
		Destination:   "a dest",
		HourlyPay:     5,
		DistanceTip:   3,
		MutualStake:   7,
	})
	require.NotNil(t, err)

	_, err = msgServer.RequestRide(context, &types.MsgRequestRide{
		Creator:       alice,
		StartLocation: "a loc",
		Destination:   "a dest",
		HourlyPay:     5,
		DistanceTip:   3,
		MutualStake:   90,
	})
	require.Nil(t, err)
}

func TestRequestRideEventEmitted(t *testing.T) {
	msgServer, _, context := setupMsgServerWithDefaultGenesis(t)
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

func TestAcceptValidation(t *testing.T) {
	msgServer, _, context := setupMsgServerAcceptRide(t)

	// Ensure that passenger cannot be driver.
	acceptRideResponse, err := msgServer.Accept(context, &types.MsgAccept{
		Creator: alice,
		IdValue: "1",
	})
	require.EqualValues(t, types.MsgAcceptResponse{
		Success: false,
	}, *acceptRideResponse)
	require.NotNil(t, err)

	acceptRideResponse, err = msgServer.Accept(context, &types.MsgAccept{
		Creator: bob,
		IdValue: "1",
	})
	// Require success in msg server handling.
	require.Nil(t, err)
	require.EqualValues(t, types.MsgAcceptResponse{
		Success: true,
	}, *acceptRideResponse)

	// Ensure that another driver cannot accept ride after it has been accepted by bob.
	acceptRideResponse, err = msgServer.Accept(context, &types.MsgAccept{
		Creator: carol,
		IdValue: "1",
	})
	require.EqualValues(t, types.MsgAcceptResponse{
		Success: false,
	}, *acceptRideResponse)
	require.NotNil(t, err)
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
		Index:            "1",
		Destination:      "some other loc",
		DriverAddress:    bob,
		PassengerAddress: alice,
		MutualStake:      50,
		PayPerHour:       15,
		DistanceTip:      10,
		AcceptanceTime:   types.TimeToString(sdk.UnwrapSDKContext(context).BlockTime()),
		BeforeId:         "-1",
		AfterId:          "-1",
		Deadline:         types.TimeToString(sdk.UnwrapSDKContext(context).BlockTime().Add(types.DeadlinePeriod)),
	}, ride1)
	require.NotEmpty(t, ride1.AcceptanceTime)
	require.NotEmpty(t, ride1.Deadline)

	deadline, err := ride1.GetDeadlineFormatted()
	require.Nil(t, err)
	acceptance, err := ride1.GetAcceptanceTimeFormatted()
	require.Nil(t, err)
	require.EqualValues(t, deadline.Sub(acceptance), 2*time.Minute)

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

func TestCreate3RidesHaveSavedFifo(t *testing.T) {

	// Request 3 rides separately.

	msgServer, keeper, context := setupMsgServerWithDefaultGenesis(t)
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

	reqResponse, err = msgServer.RequestRide(context, &types.MsgRequestRide{
		Creator:       bob,
		StartLocation: "some lat/long",
		Destination:   "some other lat/long",
		MutualStake:   50,
		HourlyPay:     25,
		DistanceTip:   10,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgRequestRideResponse{
		IdValue: "2",
	}, *reqResponse)

	reqResponse, err = msgServer.RequestRide(context, &types.MsgRequestRide{
		Creator:       carol,
		StartLocation: "some lat/long",
		Destination:   "some other lat/long",
		MutualStake:   50,
		HourlyPay:     25,
		DistanceTip:   10,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgRequestRideResponse{
		IdValue: "3",
	}, *reqResponse)

	// Now obtain those 3 rides from store and validate FIFO links.
	var tests = []struct {
		creator, idValue, beforeId, afterId string
	}{
		{alice, "1", "-1", "2"},
		{bob, "2", "1", "3"},
		{carol, "3", "2", "-1"},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%s,%s,%s,%s", tt.creator, tt.idValue, tt.beforeId, tt.afterId)
		t.Run(name, func(t *testing.T) {

			ride, found := keeper.GetStoredRide(sdk.UnwrapSDKContext(context), tt.idValue)

			if !found {
				t.Fail()
			}

			require.EqualValues(t, types.StoredRide{
				Index:            tt.idValue,
				Destination:      "some other lat/long",
				DriverAddress:    "", // Driver should not be set yet.
				PassengerAddress: tt.creator,
				MutualStake:      50,
				PayPerHour:       25,
				DistanceTip:      10,
				BeforeId:         tt.beforeId,
				AfterId:          tt.afterId,
				Deadline:         types.TimeToString(sdk.UnwrapSDKContext(context).BlockTime().Add(types.DeadlinePeriod)),
			}, ride)
		})
	}
}

func setupMsgServerFinishRide(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.RideKeeper(t)
	ride.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)
	// Request and accept ride within setup function for ease.
	response, _ := server.RequestRide(context, &types.MsgRequestRide{
		Creator:       alice,
		StartLocation: "some loc",
		Destination:   "some other loc",
		MutualStake:   50,
		HourlyPay:     15,
		DistanceTip:   10,
	})
	server.Accept(context, &types.MsgAccept{
		Creator: bob,
		IdValue: response.IdValue,
	})
	return server, *k, context
}

func TestFinishRide(t *testing.T) {
	var tests = []struct {
		creator, idValue, location string
		success                    bool
	}{
		{bob, "1", "some loc", true},
		{bob, "1", "some other loc", true},
		{alice, "1", "some loc", true},
		{alice, "1", "some other other loc", true},
		{bob, "3", "some loc", false},
		{alice, "-2", "some other loc", false},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%s,%s,%s,%t", tt.creator, tt.idValue, tt.location, tt.success)
		t.Run(name, func(t *testing.T) {
			msgServer, _, context := setupMsgServerFinishRide(t)
			finishRideResponse, err := msgServer.Finish(context, &types.MsgFinish{
				Creator:  tt.creator,
				IdValue:  tt.idValue,
				Location: tt.location,
			})
			if tt.success {
				require.Nil(t, err)
				require.EqualValues(t, types.MsgAcceptResponse{
					Success: true,
				}, *finishRideResponse)
			} else {
				require.NotNil(t, err)
				require.EqualValues(t, types.MsgAcceptResponse{
					Success: false,
				}, *finishRideResponse)
			}
		})
	}
}

func TestFinishValidation(t *testing.T) {
	msgServer, _, context := setupMsgServerFinishRide(t)

	// Ensure that an irrelevant account cannot finish ride.
	finishRideResponse, err := msgServer.Finish(context, &types.MsgFinish{
		Creator:  carol,
		IdValue:  "1",
		Location: "finish loc from rando",
	})
	require.EqualValues(t, types.MsgFinishResponse{
		Success: false,
	}, *finishRideResponse)
	require.NotNil(t, err)

	finishRideResponse, err = msgServer.Finish(context, &types.MsgFinish{
		Creator:  bob,
		IdValue:  "1",
		Location: "finish loc",
	})
	// Require success in msg server handling.
	require.Nil(t, err)
	require.EqualValues(t, types.MsgFinishResponse{
		Success: true,
	}, *finishRideResponse)

	// Ensure that this same game cannot be finished again.
	finishRideResponse, err = msgServer.Finish(context, &types.MsgFinish{
		Creator:  bob,
		IdValue:  "1",
		Location: "finish loc that's different",
	})
	require.EqualValues(t, types.MsgFinishResponse{
		Success: false,
	}, *finishRideResponse)
	require.NotNil(t, err)

	// Even by another creator
	finishRideResponse, err = msgServer.Finish(context, &types.MsgFinish{
		Creator:  carol,
		IdValue:  "1",
		Location: "finish loc that's different",
	})
	require.EqualValues(t, types.MsgFinishResponse{
		Success: false,
	}, *finishRideResponse)
	require.NotNil(t, err)
}

func TestFinishRideStorage(t *testing.T) {
	msgServer, keeper, context := setupMsgServerFinishRide(t)
	finishRideResponse, err := msgServer.Finish(context, &types.MsgFinish{
		Creator:  bob,
		IdValue:  "1",
		Location: "finish loc",
	})
	// Require success in msg server handling.
	require.Nil(t, err)
	require.EqualValues(t, types.MsgFinishResponse{
		Success: true,
	}, *finishRideResponse)
	// Ensure stored ride can be accessed from key.
	ride1, found1 := keeper.GetStoredRide(sdk.UnwrapSDKContext(context), "1")
	require.True(t, found1)
	// Ensure no other fields were mutated, and that ride is now finished.
	require.EqualValues(t, types.StoredRide{
		Index:            "1",
		Destination:      "some other loc",
		DriverAddress:    bob,
		PassengerAddress: alice,
		MutualStake:      50,
		PayPerHour:       15,
		DistanceTip:      10,
		// Block time returns a default value in unit tests, so these two timestamps will be equiv.
		AcceptanceTime: types.TimeToString(sdk.UnwrapSDKContext(context).BlockTime()),
		FinishTime:     types.TimeToString(sdk.UnwrapSDKContext(context).BlockTime()),
		FinishLocation: "finish loc",
		BeforeId:       "-1",
		AfterId:        "-1",
		Deadline:       types.TimeToString(sdk.UnwrapSDKContext(context).BlockTime().Add(types.DeadlinePeriod)),
	}, ride1)
	require.NotEmpty(t, ride1.AcceptanceTime)
	require.NotEmpty(t, ride1.FinishTime)
	require.NotEmpty(t, ride1.Deadline)

	deadline, err := ride1.GetDeadlineFormatted()
	require.Nil(t, err)
	acceptance, err := ride1.GetAcceptanceTimeFormatted()
	require.Nil(t, err)
	require.EqualValues(t, deadline.Sub(acceptance), 2*time.Minute)

	finished, err := ride1.GetFinishTimeFormatted()
	require.Nil(t, err)
	require.EqualValues(t, deadline.Sub(finished), 2*time.Minute)

	require.True(t, ride1.HasAssignedDriver())
	require.True(t, ride1.IsFinished())
}

func TestFinishRideEventEmitted(t *testing.T) {
	msgServer, _, context := setupMsgServerFinishRide(t)
	finishRideResponse, err := msgServer.Finish(context, &types.MsgFinish{
		Creator:  bob,
		IdValue:  "1",
		Location: "finish loc",
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgFinishResponse{
		Success: true,
	}, *finishRideResponse)

	ctx := sdk.UnwrapSDKContext(context)
	require.NotNil(t, ctx)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 1)
	event := events[0]
	require.EqualValues(t, []sdk.Attribute{
		{Key: "module", Value: "ride"},
		{Key: "action", Value: "RideFinished"},
		// Throw out attributes 0 through 8, coming from request/accept ride event.
	}, event.Attributes[9:])
}

func setupMsgServerRateRide(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.RideKeeper(t)
	ride.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)
	response, _ := server.RequestRide(context, &types.MsgRequestRide{
		Creator:       alice,
		StartLocation: "some loc",
		Destination:   "some other loc",
		MutualStake:   50,
		HourlyPay:     15,
		DistanceTip:   10,
	})
	server.Accept(context, &types.MsgAccept{
		Creator: bob,
		IdValue: response.IdValue,
	})
	server.Finish(context, types.NewMsgFinish(bob, "1", "some other loc"))
	return server, *k, context
}

func TestRateRide(t *testing.T) {
	msgServer, keeper, context := setupMsgServerRateRide(t)
	ctx := sdk.UnwrapSDKContext(context)
	rateRideResponse, err := msgServer.Rate(context, &types.MsgRate{
		Creator: bob,
		RideId:  "1",
		Ratee:   alice,
		Rating:  "9.5",
	})
	require.Nil(t, err)
	require.EqualValues(t, *rateRideResponse, types.MsgRateResponse{
		Success: true,
	})
	rateStruct, found := keeper.GetRatingStruct(ctx, alice)
	require.True(t, found)
	ratingAsFloat, err := strconv.ParseFloat(rateStruct.Rating, 64)
	require.Nil(t, err)
	require.EqualValues(t, 9.5*0.1, ratingAsFloat)

	// Test second rating.
	response, _ := msgServer.RequestRide(context, &types.MsgRequestRide{
		Creator:       alice,
		StartLocation: "some loc",
		Destination:   "some other loc",
		MutualStake:   50,
		HourlyPay:     15,
		DistanceTip:   10,
	})
	msgServer.Accept(context, &types.MsgAccept{
		Creator: bob,
		IdValue: response.IdValue,
	})
	msgServer.Finish(context, types.NewMsgFinish(bob, "1", "some other loc"))

	rateRideResponse, err = msgServer.Rate(context, &types.MsgRate{
		Creator: bob,
		RideId:  "1",
		Ratee:   alice,
		Rating:  "8.5",
	})
	require.Nil(t, err)
	require.EqualValues(t, *rateRideResponse, types.MsgRateResponse{
		Success: true,
	})
	rateStruct, found = keeper.GetRatingStruct(ctx, alice)
	require.True(t, found)
	ratingAsFloat, err = strconv.ParseFloat(rateStruct.Rating, 64)
	require.Nil(t, err)
	require.EqualValues(t, 9.5*0.1*0.9+8.5*0.1, ratingAsFloat)
}

func TestGasConsumption(t *testing.T) {
	msgServer, _, context := setupMsgServerWithDefaultGenesis(t)
	sdkCtx := sdk.UnwrapSDKContext(context)
	gasBefore := sdkCtx.GasMeter().GasConsumed()
	msgServer.RequestRide(context, &types.MsgRequestRide{
		Creator:       alice,
		StartLocation: "some loc",
		Destination:   "some dest",
		MutualStake:   30,
		HourlyPay:     15,
		DistanceTip:   5,
	})
	gasAfter := sdkCtx.GasMeter().GasConsumed()
	fmt.Println(gasAfter - gasBefore)
	require.EqualValues(t, uint64(0x243e)+uint64(10), gasAfter-gasBefore)
}
