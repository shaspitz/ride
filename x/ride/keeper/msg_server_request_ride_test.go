package keeper_test

import (
	"testing"

	"github.com/smarshall-spitzbart/ride/x/ride/types"
	"github.com/stretchr/testify/require"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
)

func TestRequestRide(t *testing.T) {
	msgServer, context := setupMsgServer(t)
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
		IdValue: "", // TODO: update with a proper value when updated
	}, *createResponse)
}
