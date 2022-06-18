package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
)

// Returns an error if the player has not enough funds.
// TODO: Test atomicity of this method in the context of ride request.
// TODO: Ie. make sure that if passenger doesn't have enough stake, no ride req will be stored.
func (k *Keeper) CollectDriverStake(ctx sdk.Context, storedRide *types.StoredRide) error {

	driverAddress, err := storedRide.GetDriverAddress()
	if err != nil {
		panic(err.Error())
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(
		ctx, driverAddress, types.ModuleName, sdk.NewCoins(storedRide.GetMutualStakeInCoin()))
	if err != nil {
		return sdkerrors.Wrapf(err, types.ErrMutualStake.Error(),
			"driver doesn't have enough funds to match mutual stake")
	}
	return nil
}

// Returns an error if the player has not enough funds.
// TODO: Test atomicity of this method in the context of ride request.
// TODO: Ie. make sure that if passenger doesn't have enough stake, no ride req will be stored.
func (k *Keeper) CollectPassengerStake(ctx sdk.Context, storedRide *types.StoredRide) error {

	passengerAddress, err := storedRide.GetPassengerAddress()
	if err != nil {
		panic(err.Error())
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(
		ctx, passengerAddress, types.ModuleName, sdk.NewCoins(storedRide.GetMutualStakeInCoin()))
	if err != nil {
		return sdkerrors.Wrapf(err, types.ErrMutualStake.Error(),
			"passenger doesn't have enough funds to match mutual stake")
	}
	return nil
}

// at the end of a ride completing, one of these three methods will run
// TODO: Use invariant module for these "Must" methods

func (k *Keeper) MustPayout(ctx sdk.Context, storedRide *types.StoredRide) {
	// NOTE: Assume rides will be under 1 hour + distance tip for now, can protect against edge
	// cases if there's time.

	// TODO: Ensure invariant that driver and passenger addresses are populated, also times

	passengerAddress, err := storedRide.GetPassengerAddress()
	if err != nil {
		panic(err.Error())
	}

	driverAddress, err := storedRide.GetDriverAddress()
	if err != nil {
		panic(err.Error())
	}

	finishTime, err := storedRide.GetFinishTimeFormatted()
	if err != nil {
		panic(err.Error())
	}

	acceptanceTime, err := storedRide.GetAcceptanceTimeFormatted()
	if err != nil {
		panic(err.Error())
	}
	driveDuration := finishTime.Sub(acceptanceTime)

	// Driver is paid back original stake, plus tip, and hourly pay.
	driverPayoutAmount := storedRide.MutualStake + storedRide.DistanceTip + storedRide.PayPerHour*uint64(driveDuration.Hours())
	driverPayoutInCoin := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(driverPayoutAmount)))

	// Passenger is paid back remaining funds from ride account.
	passengerPayout := 2*storedRide.MutualStake - driverPayoutAmount
	passengerPayoutInCoin := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(passengerPayout)))

	// TODO: Enforce invariant that passenger payout is still positive.

	err = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, driverAddress, sdk.NewCoins(driverPayoutInCoin))
	if err != nil {
		panic(err.Error())
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, passengerAddress, sdk.NewCoins(passengerPayoutInCoin))
	if err != nil {
		panic(err.Error())
	}
}

func (k *Keeper) MustProcessDispute(ctx sdk.Context, storedRide *types.StoredRide) {
	// TODO, implement disputes later on.
}

// TODO: Unit test.
func (k *Keeper) MustRefundStakes(ctx sdk.Context, storedRide *types.StoredRide) {

	// TODO: Ensure invariant that at least passenger address is populated.

	passengerAddress, err := storedRide.GetPassengerAddress()
	if err != nil {
		panic(err.Error())
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, passengerAddress, sdk.NewCoins(storedRide.GetMutualStakeInCoin()))
	if err != nil {
		panic(err.Error())
	}

	// No one else to refund if driver has not been assigned to ride.
	if !storedRide.HasAssignedDriver() {
		return
	}

	driverAddress, err := storedRide.GetDriverAddress()
	if err != nil {
		panic(err.Error())
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, driverAddress, sdk.NewCoins(storedRide.GetMutualStakeInCoin()))
	if err != nil {
		panic(err.Error())
	}
}
