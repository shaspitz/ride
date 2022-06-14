package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
)

func (k msgServer) Accept(goCtx context.Context, msg *types.MsgAccept) (*types.MsgAcceptResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storedRide, found := k.Keeper.GetStoredRide(ctx, msg.IdValue)
	if !found {
		return &types.MsgAcceptResponse{Success: false},
			errors.Wrapf(types.ErrRideNotFound, "ride not found %s", msg.IdValue)
	}

	err := ValidateAcceptRide(msg, storedRide)
	if err != nil {
		return &types.MsgAcceptResponse{Success: false}, err
	}

	// Assign driver to ride.
	storedRide.Driver = msg.Creator

	// Validate assigned driver address, could be in unit test to save gas.
	_, err = storedRide.GetDriverAddress()
	if err != nil {
		return &types.MsgAcceptResponse{Success: false},
			errors.Wrapf(types.ErrInvalidDriver, "invalid driver address %s", msg.IdValue)
	}

	// Store acceptance time in default format.
	storedRide.AcceptanceTime = types.TimeToString(ctx.BlockTime())

	// TODO: Assign mutual stake from driver and passenger to the bank keeper,
	// write tests to see if this whole method runs atomically..
	// Ie. if an error is returned, does the state just get thrown out by
	// any validator who executes the msg?

	// TODO: Set expiration time for the actual ride?? and store in state??

	// Store ride via keeper.
	k.Keeper.SetStoredRide(ctx, storedRide)

	// Emit appropriate event for ride acceptance.
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AcceptRideEventKey),
			sdk.NewAttribute(types.AcceptRideEventDriver, storedRide.Driver),
			sdk.NewAttribute(types.AcceptRideEventIdValue, storedRide.Index),
		),
	)

	return &types.MsgAcceptResponse{Success: true}, nil
}

// Validation of accept message handler, in its own method
// per https://docs.cosmos.network/main/building-modules/msg-services.html#validation
//
// NOTE: Oracle validation would exist here for starting location.

func ValidateAcceptRide(msg *types.MsgAccept, storedRide types.StoredRide) error {

	if storedRide.HasAssignedDriver() {
		return errors.Wrapf(types.ErrAssignedDriver, "driver has already been assigned for ride with Id: %s", msg.IdValue)
	}
	// TODO: Enforce that driver actually possesses mutual stake set by passenger.
	// TODO: Charge any gas here?
	return nil
}
