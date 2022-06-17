// Helper methods involving a stored ride, that're not auto generated.
package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errors "github.com/cosmos/cosmos-sdk/types/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/types"
)

func (storedRide StoredRide) GetDriverAddress() (driver sdk.AccAddress, err error) {
	if !storedRide.HasAssignedDriver() {
		return nil, errors.Wrapf(err, ErrNoAssignedDriver.Error(), storedRide.Driver)
	}
	driver, err = sdk.AccAddressFromBech32(storedRide.Driver)
	return driver, errors.Wrapf(err, ErrInvalidDriver.Error(), storedRide.Driver)
}

func (storedRide StoredRide) GetPassengerAddress() (passenger sdk.AccAddress, err error) {
	passenger, err = sdk.AccAddressFromBech32(storedRide.Passenger)
	return passenger, errors.Wrapf(err, ErrInvalidPassenger.Error(), storedRide.Passenger)
}

func (storedRide StoredRide) HasAssignedDriver() bool {
	return storedRide.Driver != ""
}

func (storedRide StoredRide) IsFinished() bool {
	return storedRide.FinishTime != ""
}

func (storedRide StoredRide) HasExpired(ctx sdk.Context) (bool, error) {
	deadlineAsTime, err := storedRide.GetDeadlineFormatted()
	if err != nil {
		return false, err
	}
	return deadlineAsTime.Before(ctx.BlockTime()), nil
}

func (storedRide *StoredRide) GetAcceptanceTimeFormatted() (accepted time.Time, err error) {
	accepted, err = time.Parse(types.TimeFormat, storedRide.AcceptanceTime)
	return accepted, sdkerrors.Wrapf(err, ErrCantParseTime.Error(), storedRide.AcceptanceTime)
}

func (storedRide *StoredRide) GetFinishTimeFormatted() (finished time.Time, err error) {
	finished, err = time.Parse(types.TimeFormat, storedRide.FinishTime)
	return finished, sdkerrors.Wrapf(err, ErrCantParseTime.Error(), storedRide.FinishTime)
}

func (storedRide *StoredRide) GetDeadlineFormatted() (deadline time.Time, err error) {
	deadline, err = time.Parse(types.TimeFormat, storedRide.Deadline)
	return deadline, sdkerrors.Wrapf(err, ErrCantParseTime.Error(), storedRide.FinishTime)
}

func TimeToString(time time.Time) string {
	return time.UTC().Format(types.TimeFormat)
}
