// Helper methods involving a stored ride, that're not auto generated.
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	errors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (storedRide StoredRide) GetDriverAddress() (driver sdk.AccAddress, err error) {
	if !storedRide.HasAssignedDriver() {
		return nil, errors.Wrapf(err, ErrNoAssignedDriver.Error(), storedRide.Driver)
	}
	driver, err = sdk.AccAddressFromBech32(storedRide.Driver)
	return driver, errors.Wrapf(err, ErrInvalidDriver.Error(), storedRide.Driver)
}

func (storedRide StoredRide) HasAssignedDriver() bool {
	return storedRide.Driver != ""
}

func (storedRide StoredRide) GetPassengerAddress() (passenger sdk.AccAddress, err error) {
	passenger, err = sdk.AccAddressFromBech32(storedRide.Passenger)
	return passenger, errors.Wrapf(err, ErrInvalidPassenger.Error(), storedRide.Passenger)
}
