package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRequestRide = "request_ride"

var _ sdk.Msg = &MsgRequestRide{}

func NewMsgRequestRide(creator string, startLocation string, destination string, mutualStake uint64, hourlyPay uint64, distanceTip uint64) *MsgRequestRide {
	return &MsgRequestRide{
		Creator:       creator,
		StartLocation: startLocation,
		Destination:   destination,
		MutualStake:   mutualStake,
		HourlyPay:     hourlyPay,
		DistanceTip:   distanceTip,
	}
}

func (msg *MsgRequestRide) Route() string {
	return RouterKey
}

func (msg *MsgRequestRide) Type() string {
	return TypeMsgRequestRide
}

func (msg *MsgRequestRide) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestRide) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestRide) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
