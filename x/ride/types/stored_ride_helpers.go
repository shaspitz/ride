// Helper methods involving a stored ride, that're not auto generated.
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (storedRide StoredRide) GetDriverAddress() (driver sdk.AccAddress, err error) {
	driver, err = sdk.AccAddressFromBech32(storedRide.Driver)
	return driver, sdkerrors.Wrapf(err, ErrInvalidDriver.Error(), storedRide.Driver)
}

func (storedRide StoredRide) GetPassengerAddress() (passenger sdk.AccAddress, err error) {
	passenger, err = sdk.AccAddressFromBech32(storedRide.Passenger)
	return passenger, sdkerrors.Wrapf(err, ErrInvalidPassenger.Error(), storedRide.Passenger)
}
