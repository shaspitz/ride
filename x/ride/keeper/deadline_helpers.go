package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
)

// Helpers to implement ride deadlines more efficiently.
// See: https://tutorials.cosmos.network/academy/3-my-own-chain/ride-fifo.html

func GetNextDeadline(ctx sdk.Context) time.Time {
	return ctx.BlockTime().Add(types.DeadlinePeriod)
}

// Removes a stored ride from the FIFO doubly linked list.
// UPDATE STORE AFTER USING THIS METHOD.
func (k Keeper) RemoveFromFifo(ctx sdk.Context, ride *types.StoredRide, info *types.NextRide) {
	// Does it have a predecessor?
	if ride.BeforeId != types.NoFifoIdKey {
		beforeElement, found := k.GetStoredRide(ctx, ride.BeforeId)
		if !found {
			panic("Element before in Fifo was not found")
		}
		beforeElement.AfterId = ride.AfterId
		k.SetStoredRide(ctx, beforeElement)
		if ride.AfterId == types.NoFifoIdKey {
			info.FifoTail = beforeElement.Index
		}
		// Is it at the FIFO head?
	} else if info.FifoHead == ride.Index {
		info.FifoHead = ride.AfterId
	}
	// Does it have a successor?
	if ride.AfterId != types.NoFifoIdKey {
		afterElement, found := k.GetStoredRide(ctx, ride.AfterId)
		if !found {
			panic("Element after in Fifo was not found")
		}
		afterElement.BeforeId = ride.BeforeId
		k.SetStoredRide(ctx, afterElement)
		if ride.BeforeId == types.NoFifoIdKey {
			info.FifoHead = afterElement.Index
		}
		// Is it at the FIFO tail?
	} else if info.FifoTail == ride.Index {
		info.FifoTail = ride.BeforeId
	}
	ride.BeforeId = types.NoFifoIdKey
	ride.AfterId = types.NoFifoIdKey
}

// Appends a stored ride to the tail of the FIFO doubly linked list, removes it from previous index if needed.
// UPDATE STORE AFTER USING THIS METHOD.
func (k Keeper) SendToFifoTail(ctx sdk.Context, ride *types.StoredRide, info *types.NextRide) {

	// Update deadline before moving within linked list.
	ride.Deadline = types.TimeToString(GetNextDeadline(ctx))

	if info.FifoHead == types.NoFifoIdKey && info.FifoTail == types.NoFifoIdKey {
		ride.BeforeId = types.NoFifoIdKey
		ride.AfterId = types.NoFifoIdKey
		info.FifoHead = ride.Index
		info.FifoTail = ride.Index
	} else if info.FifoHead == types.NoFifoIdKey || info.FifoTail == types.NoFifoIdKey {
		panic("Fifo should have both head and tail or none")
	} else if info.FifoTail == ride.Index {
		// Nothing to do, already at tail
	} else {
		// Snip ride out
		k.RemoveFromFifo(ctx, ride, info)

		// Now add to tail
		currentTail, found := k.GetStoredRide(ctx, info.FifoTail)
		if !found {
			panic("Current Fifo tail was not found")
		}
		currentTail.AfterId = ride.Index
		k.SetStoredRide(ctx, currentTail)

		ride.BeforeId = currentTail.Index
		info.FifoTail = ride.Index
	}
}
