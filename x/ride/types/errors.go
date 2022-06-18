package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/ride module sentinel errors
var (
	ErrSample              = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidDriver       = sdkerrors.Register(ModuleName, 1101, "driver address is invalid")
	ErrNoAssignedDriver    = sdkerrors.Register(ModuleName, 1102, "driver has not been assigned for this ride")
	ErrAssignedDriver      = sdkerrors.Register(ModuleName, 1103, "driver has already been assigned for this ride")
	ErrInvalidPassenger    = sdkerrors.Register(ModuleName, 1104, "passenger address is invalid")
	ErrRideNotFound        = sdkerrors.Register(ModuleName, 1105, "ride by id not found, it may have expired")
	ErrCantParseTime       = sdkerrors.Register(ModuleName, 1106, "cant parse time from string %s")
	ErrIrrelevantRide      = sdkerrors.Register(ModuleName, 1107, "msg creator not passenger or driver for this ride")
	ErrAlreadyFinishedRide = sdkerrors.Register(ModuleName, 1108, "ride already finished")
	ErrCannotDriveYourself = sdkerrors.Register(ModuleName, 1109, "passenger cannot be driver")

	ErrMutualStake     = sdkerrors.Register(ModuleName, 1110, "error processing mutual stake")
	ErrCantRefundStake = sdkerrors.Register(ModuleName, 1111, "cannot refund stake")
	ErrRideParameters  = sdkerrors.Register(ModuleName, 1112, "ride parameters are invalid")
)
