package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/ride module sentinel errors
var (
	ErrSample           = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidDriver    = sdkerrors.Register(ModuleName, 1101, "driver address is invalid: %s")
	ErrInvalidPassenger = sdkerrors.Register(ModuleName, 1102, "passenger address is invalid: %s")
)
