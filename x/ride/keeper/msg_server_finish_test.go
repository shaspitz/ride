package keeper_test

import (
	"context"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/smarshall-spitzbart/ride/testutil/keeper"
	"github.com/smarshall-spitzbart/ride/x/ride"
	"github.com/smarshall-spitzbart/ride/x/ride/keeper"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
	"github.com/stretchr/testify/require"
)

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
		Index:          "1",
		Destination:    "some other loc",
		Driver:         bob,
		Passenger:      alice,
		MutualStake:    50,
		PayPerHour:     15,
		DistanceTip:    10,
		AcceptanceTime: ride1.AcceptanceTime,
		FinishTime:     ride1.FinishTime,
		FinishLocation: "finish loc",
	}, ride1)
	require.NotEmpty(t, ride1.AcceptanceTime)
	require.NotEmpty(t, ride1.FinishTime)

	_, err = ride1.GetFinishTimeFormatted()
	require.Nil(t, err)
	_, err = ride1.GetAcceptanceTimeFormatted()
	require.Nil(t, err)

	// TODO: Properly mock block time passing by.
	// require.Positive(t, finishTime.Sub(acceptanceTime))

	require.True(t, ride1.HasAssignedDriver())
	require.True(t, ride1.IsFinished())
}

// TODO: Make tests for the functionality that'll be included in "ValidateFinishRide"

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
