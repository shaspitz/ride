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
	game1, found1 := keeper.GetStoredRide(sdk.UnwrapSDKContext(context), "1")
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
	}, game1)
	// Ensure driver is assigned via appropriate method.
	require.True(t, game1.HasAssignedDriver())
	// Finally, ensure that another driver cannot accept the ride.
	acceptRideResponse, err = msgServer.Accept(context, &types.MsgAccept{
		Creator: carol,
	})
	require.NotNil(t, err)
	require.EqualValues(t, types.MsgAcceptResponse{
		Success: false,
	}, *acceptRideResponse)
}

// TODO: Make tests for the functionality that'll be included in "ValidateRequestRide"
