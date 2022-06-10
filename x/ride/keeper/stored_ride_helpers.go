// Helper methods involving a stored ride, that're not auto generated.
package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
)

func (k Keeper) GetDriverAccount(storedRide types.StoredRide) (driver sdk.AccAddress, err error) {
	driver, err = sdk.AccAddressFromBech32(storedRide.Driver)
	return driver, sdkerrors.Wrapf(err, "Driver account address could not be obtained from string")
}

func (k Keeper) GetPassengerAccount(storedRide types.StoredRide) (passenger sdk.AccAddress, err error) {
	passenger, err = sdk.AccAddressFromBech32(storedRide.Driver)
	return passenger, sdkerrors.Wrapf(err, "Passenger account address could not be obtained from string")
}