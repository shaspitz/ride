package keeper_test

import (
	"context"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/smarshall-spitzbart/ride/testutil/keeper"
	"github.com/smarshall-spitzbart/ride/x/ride"
	"github.com/smarshall-spitzbart/ride/x/ride/keeper"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
	"github.com/stretchr/testify/require"
)

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
