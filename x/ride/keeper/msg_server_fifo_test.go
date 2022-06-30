package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
	"github.com/stretchr/testify/require"
)

// TODO: Add more tests inspired from https://tutorials.cosmos.network/academy/3-my-own-chain/game-fifo.html#unit-tests

func TestCreate3RidesHaveSavedFifo(t *testing.T) {

	// Request 3 rides separately.

	msgServer, keeper, context := setupMsgServerRequestRide(t)
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
