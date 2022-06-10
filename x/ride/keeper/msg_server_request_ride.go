package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
)

func (k msgServer) RequestRide(goCtx context.Context, msg *types.MsgRequestRide) (*types.MsgRequestRideResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRequestRideResponse{}, nil
}
