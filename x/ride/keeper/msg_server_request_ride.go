package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errors "github.com/cosmos/cosmos-sdk/types/errors"
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
		// To be updated below.
		BeforeId: types.NoFifoIdKey,
		AfterId:  types.NoFifoIdKey,
	}

	// Validate assigned passenger address can be parsed.
	_, err = storedRide.GetPassengerAddress()
	if err != nil {
		return nil, err
	}

	err = k.Keeper.CollectDriverStake(ctx, &storedRide)
	if err != nil {
		return nil, err
	}

	// Ride is stored this block, send it to the FIFO tail.
	k.Keeper.SendToFifoTail(ctx, &storedRide, &nextRide)

	// Store ride via keeper.
	k.Keeper.SetStoredRide(ctx, storedRide)

	// Increment counter for next stored ride.
	nextRide.IdValue++
	k.Keeper.SetNextRide(ctx, nextRide)

	// Emit appropriate event with starting location of ride.
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.RequestRideEventKey),
			sdk.NewAttribute(types.RequestRideEventPassenger, storedRide.Passenger),
			sdk.NewAttribute(types.RequestRideEventIndex, storedRide.Index),
			sdk.NewAttribute(types.RequestRideStartLocation, msg.StartLocation),
		),
	)

	return &types.MsgRequestRideResponse{
		IdValue: newRideIndex,
	}, nil
}

// Validation of RequestRide message handler, in its own method
// per https://docs.cosmos.network/main/building-modules/msg-services.html#validation
//
// NOTE: Oracle validation would exist here for starting location.
func ValidateRequestRide(msg *types.MsgRequestRide) error {

	if msg.MutualStake < msg.DistanceTip+msg.HourlyPay {
		return errors.Wrapf(types.ErrRideParameters, "mutual stake is below distance tip + 1 hour pay")
	}

	// TODO: Enforce that passenger actually possesses mutual stake set by Tx <- Auth module?

	// TODO: Charge any gas here?
	return nil
}
