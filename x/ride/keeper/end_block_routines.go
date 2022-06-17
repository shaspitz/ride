package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
)

type ExpiredRide interface {
	HandleExpiration(k Keeper) bool
	InvokeEvent(k Keeper) bool
}

// Going back in time, it'd be better to implement these three structures as their own
// objects stored to disk, not wrappers around a single ride object with a lot of fields.

type ExpiredRideRequest struct {
	storedRide *types.StoredRide
}

type ExpiredActiveRide struct {
	storedRide *types.StoredRide
}

type ExpiredFinishedRide struct {
	storedRide *types.StoredRide
}

func (k Keeper) CleanupExpiredRides(goCtx context.Context) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	nextRide, found := k.GetNextRide(ctx)
	if !found {
		panic("NextRide not found")
	}
	oldestStoredRideId := nextRide.FifoHead

	for {
		oldestStoredRide, found := k.GetStoredRide(ctx, oldestStoredRideId)
		if !found {
			panic("Fifo head game not found " + oldestStoredRideId)
		}
		expired, err := oldestStoredRide.HasExpired(ctx)
		if err != nil {
			panic("game expiration could not be obtained " + err.Error())
		}
		if expired {
			// for any expired ride implementation: instantiate, handle
			// (including garbage collect), then invoke event.
			var expiredRide ExpiredRide

			if oldestStoredRide.IsFinished() {
				if oldestStoredRide.HasAssignedDriver() {
					expiredRide = &ExpiredFinishedRide{storedRide: &oldestStoredRide}
				} else {
					// TODO, implement this case in finish msg handler.
					panic("finished ride without driver should have already been handled and cleaned up.")
				}
			} else {
				if oldestStoredRide.HasAssignedDriver() {
					expiredRide = &ExpiredActiveRide{storedRide: &oldestStoredRide}
				} else {
					expiredRide = &ExpiredRideRequest{storedRide: &oldestStoredRide}
				}
			}

			expiredRide.HandleExpiration(k)
			expiredRide.InvokeEvent(k)
			oldestStoredRideId = nextRide.FifoHead
		} else {
			// Since games are stored in FIFO, there are no more to check for expiration.
			break
		}
	}
	// Ensure mutations to FIFO are persisted.
	k.SetNextRide(ctx, nextRide)
}

// TODO: unit tests on all of these.

func (expRide *ExpiredRideRequest) HandleExpiration(
	k Keeper, ctx sdk.Context, nextRide *types.NextRide) {
	// TODO: Return staked funds to passenger.
	k.RemoveFromFifo(ctx, expRide.storedRide, nextRide)
	k.RemoveStoredRide(ctx, expRide.storedRide.Index)
}

func (expRide *ExpiredActiveRide) HandleExpiration(
	k Keeper, ctx sdk.Context, nextRide *types.NextRide) {
	// Convert to finished ride, dispute will likely occur.
	expRide.storedRide.FinishTime = types.TimeToString(ctx.BlockTime())
	expRide.storedRide.FinishLocation = "unknown"
	k.SendToFifoTail(ctx, expRide.storedRide, nextRide)
	k.SetStoredRide(ctx, *expRide.storedRide)
}

func (expRide *ExpiredFinishedRide) HandleExpiration(
	k Keeper, ctx sdk.Context, nextRide *types.NextRide) {
	// TODO: Resolve payouts for nominally finished rides, then garbage collect.
	k.RemoveFromFifo(ctx, expRide.storedRide, nextRide)
	k.RemoveStoredRide(ctx, expRide.storedRide.Index)
}

func (expRide *ExpiredRideRequest) InvokeEvent(k Keeper, ctx sdk.Context) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.RideRequestExpiredEventKey),
		),
	)
}

func (expRide *ExpiredActiveRide) InvokeEvent(k Keeper, ctx sdk.Context) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.ActiveRideExpiredEventKey),
		),
	)
}

func (expRide *ExpiredFinishedRide) InvokeEvent(k Keeper, ctx sdk.Context) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.FinishedRideExpiredEventKey),
		),
	)
}
