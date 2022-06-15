package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
)

// Helpers to implement ride deadlines more efficiently.
// See: https://tutorials.cosmos.network/academy/3-my-own-chain/game-fifo.html

// TODO: Make this hardcoded value configurable.
func GetNextDeadline(ctx sdk.Context) time.Time {
	return ctx.BlockTime().Add(5 * time.Minute)
}

// UPDATE STORE AFTER USING THIS METHOD.
func (k Keeper) RemoveFromFifo(ctx sdk.Context, game *types.StoredRide, info *types.NextRide) {
	// Does it have a predecessor?
	if game.BeforeId != types.NoFifoIdKey {
		beforeElement, found := k.GetStoredRide(ctx, game.BeforeId)
		if !found {
			panic("Element before in Fifo was not found")
		}
		beforeElement.AfterId = game.AfterId
		k.SetStoredRide(ctx, beforeElement)
		if game.AfterId == types.NoFifoIdKey {
			info.FifoTail = beforeElement.Index
		}
		// Is it at the FIFO head?
	} else if info.FifoHead == game.Index {
		info.FifoHead = game.AfterId
	}
	// Does it have a successor?
	if game.AfterId != types.NoFifoIdKey {
		afterElement, found := k.GetStoredRide(ctx, game.AfterId)
		if !found {
			panic("Element after in Fifo was not found")
		}
		afterElement.BeforeId = game.BeforeId
		k.SetStoredRide(ctx, afterElement)
		if game.BeforeId == types.NoFifoIdKey {
			info.FifoHead = afterElement.Index
		}
		// Is it at the FIFO tail?
	} else if info.FifoTail == game.Index {
		info.FifoTail = game.BeforeId
	}
	game.BeforeId = types.NoFifoIdKey
	game.AfterId = types.NoFifoIdKey
}

// UPDATE STORE AFTER USING THIS METHOD.
func (k Keeper) SendToFifoTail(ctx sdk.Context, game *types.StoredRide, info *types.NextRide) {
	if info.FifoHead == types.NoFifoIdKey && info.FifoTail == types.NoFifoIdKey {
		game.BeforeId = types.NoFifoIdKey
		game.AfterId = types.NoFifoIdKey
		info.FifoHead = game.Index
		info.FifoTail = game.Index
	} else if info.FifoHead == types.NoFifoIdKey || info.FifoTail == types.NoFifoIdKey {
		panic("Fifo should have both head and tail or none")
	} else if info.FifoTail == game.Index {
		// Nothing to do, already at tail
	} else {
		// Snip game out
		k.RemoveFromFifo(ctx, game, info)

		// Now add to tail
		currentTail, found := k.GetStoredRide(ctx, info.FifoTail)
		if !found {
			panic("Current Fifo tail was not found")
		}
		currentTail.AfterId = game.Index
		k.SetStoredRide(ctx, currentTail)

		game.BeforeId = currentTail.Index
		info.FifoTail = game.Index
	}
}
