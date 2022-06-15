package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
)

func (k msgServer) Finish(goCtx context.Context, msg *types.MsgFinish) (*types.MsgFinishResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storedRide, found := k.Keeper.GetStoredRide(ctx, msg.IdValue)
	if !found {
		return &types.MsgFinishResponse{Success: false},
			errors.Wrapf(types.ErrRideNotFound, "ride not found with Id: %s", msg.IdValue)
	}

	nextRide, found := k.Keeper.GetNextRide(ctx)
	if !found {
		panic("NextRide not found")
	}

	err := ValidateFinishRide(msg, storedRide)
	if err != nil {
		return nil, err
	}

	if !storedRide.HasAssignedDriver() {
		// TODO: No driver accepted yet in this case. Return funds and erase game.
		_ = ctx
		k.Keeper.RemoveStoredRide(ctx, msg.IdValue)
	}

	// TODO: Auto execution timer? Where payments are made after a timeout!!

	storedRide.FinishTime = types.TimeToString(ctx.BlockTime())
	storedRide.FinishLocation = msg.Location

	// Ride is mutated this block, send it to the FIFO tail then update store accordingly.
	k.Keeper.SendToFifoTail(ctx, &storedRide, &nextRide)
	k.Keeper.SetStoredRide(ctx, storedRide)
	k.Keeper.SetNextRide(ctx, nextRide)

	// Emit appropriate event for ride acceptance.
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.FinishRideEventKey),
		),
	)

	return &types.MsgFinishResponse{Success: true}, nil
}

// TODO: Unit test
func ValidateFinishRide(msg *types.MsgFinish, storedRide types.StoredRide) error {
	if msg.Creator != storedRide.Passenger && msg.Creator != storedRide.Driver {
		return errors.Wrapf(types.ErrIrrelevantRide, "%s is not associated with game %s",
			msg.Creator, msg.IdValue)
	}
	return nil
}
