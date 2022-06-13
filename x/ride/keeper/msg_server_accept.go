package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
)

func (k msgServer) Accept(goCtx context.Context, msg *types.MsgAccept) (*types.MsgAcceptResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := ValidateAcceptRide(msg)
	if err != nil {
		return nil, err
	}

	storedRide, found := k.Keeper.GetStoredRide(ctx, msg.IdValue)
	if !found {
		return &types.MsgAcceptResponse{Success: false},
			errors.Wrapf(types.ErrRideNotFound, "game not found %s", msg.IdValue)
	}

	if storedRide.HasAssignedDriver() {
		return &types.MsgAcceptResponse{Success: false},
			errors.Wrapf(types.ErrAssignedDriver, "driver has already been assigned for this ride %s", msg.IdValue)
	}

	// Assign driver to ride.
	storedRide.Driver = msg.Creator

	// Validate assigned driver address.
	_, err = storedRide.GetDriverAddress()
	if err != nil {
		return &types.MsgAcceptResponse{Success: false},
			errors.Wrapf(types.ErrInvalidDriver, "invalid driver address %s", msg.IdValue)
	}

	// TODO: send driver's tokens to escrow, make a test to ensure that this method runs atomically.

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
func ValidateAcceptRide(msg *types.MsgAccept) error {

	// TODO: Enforce that driver actually possesses mutual stake set by passenger.
	// TODO: Charge any gas here?
	// TODO: Set expiration time for the actual ride?? and store in state??
	return nil
}
