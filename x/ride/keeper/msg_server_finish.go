package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
)

func (k msgServer) Finish(goCtx context.Context, msg *types.MsgFinish) (*types.MsgFinishResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := ValidateFinishRide(msg)
	if err != nil {
		return nil, err
	}

	storedRide, found := k.Keeper.GetStoredRide(ctx, msg.IdValue)
	if !found {
		return &types.MsgFinishResponse{Success: false},
			errors.Wrapf(types.ErrRideNotFound, "ride not found with Id: %s", msg.IdValue)
	}

	// TODO: Auto execution timer?

	storedRide.FinishTime = types.TimeToString(ctx.BlockTime())
	storedRide.FinishLocation = msg.Location

	// Store ride via keeper.
	k.Keeper.SetStoredRide(ctx, storedRide)

	// Emit appropriate event for ride acceptance.
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.FinishRideEventKey),
		),
	)

	return &types.MsgFinishResponse{Success: true}, nil
}

// TODO: Populate this
func ValidateFinishRide(msg *types.MsgFinish) error {

	return nil
}
