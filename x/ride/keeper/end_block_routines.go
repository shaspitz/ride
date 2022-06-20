package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
)

type ExpiredRide interface {
	HandleExpiration(k Keeper, ctx sdk.Context, nextRide *types.NextRide)
	InvokeEvent(ctx sdk.Context)
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
		// Check for being finished moving along FIFO.
		if strings.Compare(oldestStoredRideId, types.NoFifoIdKey) == 0 {
			break
		}
		oldestStoredRide, found := k.GetStoredRide(ctx, oldestStoredRideId)
		if !found {
			panic("Fifo head ride not found " + oldestStoredRideId)
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
					panic("finished ride without driver should have already been handled and cleaned up.")
				}
			} else {
				if oldestStoredRide.HasAssignedDriver() {
					expiredRide = &ExpiredActiveRide{storedRide: &oldestStoredRide}
				} else {
					expiredRide = &ExpiredRideRequest{storedRide: &oldestStoredRide}
				}
			}

			expiredRide.HandleExpiration(k, ctx, &nextRide)
			expiredRide.InvokeEvent(ctx)
			oldestStoredRideId = nextRide.FifoHead
		} else {
			// Since games are stored in FIFO, there are no more to check for expiration.
			break
		}
	}
	// Ensure mutations to FIFO are persisted.
	k.SetNextRide(ctx, nextRide)
}

func (expRide *ExpiredRideRequest) HandleExpiration(
	k Keeper, ctx sdk.Context, nextRide *types.NextRide) {
	k.MustRefundStakes(ctx, expRide.storedRide)
	k.RemoveFromFifo(ctx, expRide.storedRide, nextRide)
	k.RemoveStoredRide(ctx, expRide.storedRide.Index)
}

func (expRide *ExpiredActiveRide) HandleExpiration(
	k Keeper, ctx sdk.Context, nextRide *types.NextRide) {
	// Convert to finished ride, dispute should occur.
	expRide.storedRide.FinishTime = types.TimeToString(ctx.BlockTime())
	expRide.storedRide.FinishLocation = "unknown"
	k.SendToFifoTail(ctx, expRide.storedRide, nextRide)
	k.SetStoredRide(ctx, *expRide.storedRide)
}

func (expRide *ExpiredFinishedRide) HandleExpiration(
	k Keeper, ctx sdk.Context, nextRide *types.NextRide) {
	k.MustPayout(ctx, expRide.storedRide)
	k.RemoveFromFifo(ctx, expRide.storedRide, nextRide)
	k.RemoveStoredRide(ctx, expRide.storedRide.Index)
}

func (expRide *ExpiredRideRequest) InvokeEvent(ctx sdk.Context) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.RideRequestExpiredEventKey),
		),
	)
}

func (expRide *ExpiredActiveRide) InvokeEvent(ctx sdk.Context) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.ActiveRideExpiredEventKey),
		),
	)
}

func (expRide *ExpiredFinishedRide) InvokeEvent(ctx sdk.Context) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.FinishedRideExpiredEventKey),
		),
	)
}
