package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
)

// RequestRide message handler obtains the next ride counter from app chain state,
// creates and stores a new ride object using this counter, increments that counter
// for future rides, and returns a response.
func (k msgServer) RequestRide(goCtx context.Context, msg *types.MsgRequestRide) (*types.MsgRequestRideResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := ValidateRequestRide(msg)
	if err != nil {
		return nil, err
	}

	nextRide, found := k.Keeper.GetNextRide(ctx)
	if !found {
		panic("NextRide not found")
	}

	newRideIndex := strconv.FormatUint(nextRide.IdValue, 10)

	storedRide := types.StoredRide{
		Index: newRideIndex,
		// NOTE: Start location would be validated/stored if a proof-of-location system were implemented.
		Destination: msg.Destination,
		// NOTE: Driver is set in the accept Tx.
		Passenger:   msg.Creator,
		MutualStake: msg.MutualStake,
		PayPerHour:  msg.HourlyPay,
		DistanceTip: msg.DistanceTip,
	}

	// Validate assigned passenger address.
	_, err = storedRide.GetPassengerAddress()
	if err != nil {
		return nil, err
	}

	// TODO: Assign mutual stake to the bank keeper, write tests to see if this whole method runs atomically..
	// Ie. if an error is returned, does the state just get thrown out by the any validator who executes the msg?

	// Store ride via keeper.
	k.Keeper.SetStoredRide(ctx, storedRide)

	// Increment counter for next stored ride.
	nextRide.IdValue++
	k.Keeper.SetNextRide(ctx, nextRide)

	return &types.MsgRequestRideResponse{
		IdValue: newRideIndex,
	}, nil
}

// Validation of RequestRide message handler, in its own method
// per https://docs.cosmos.network/main/building-modules/msg-services.html#validation
//
// NOTE: Oracle validation would exist here for starting location.
func ValidateRequestRide(msg *types.MsgRequestRide) error {

	// TODO: Enforce mutual stake is above some config, etc.
	// TODO: Enforce that passenger actually possesses mutual stake set by Tx.
	// TODO: Charge any gas here?
	// TODO: Set expiration time for the request?? and store in state??
	return nil
}
