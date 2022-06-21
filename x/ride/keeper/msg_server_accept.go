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

	nextRide, found := k.Keeper.GetNextRide(ctx)
	if !found {
		panic("NextRide not found")
	}

	err := ValidateAcceptRide(msg, storedRide)
	if err != nil {
		return &types.MsgAcceptResponse{Success: false}, err
	}

	ctx.GasMeter().ConsumeGas(types.AcceptRideGas, "Accept Ride")

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

	err = k.Keeper.CollectDriverStake(ctx, &storedRide)
	if err != nil {
		return nil, err
	}

	// Ride is mutated this block, send it to the FIFO tail then update store accordingly.
	k.Keeper.SendToFifoTail(ctx, &storedRide, &nextRide)
	k.Keeper.SetStoredRide(ctx, storedRide)
	k.Keeper.SetNextRide(ctx, nextRide)

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

	if msg.Creator == storedRide.Passenger {
		return errors.Wrapf(types.ErrCannotDriveYourself,
			"%s is the passenger of this ride, and cannot be the driver", storedRide.Passenger)
	}
	return nil
}
