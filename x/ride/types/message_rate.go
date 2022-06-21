package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRate = "rate"

var _ sdk.Msg = &MsgRate{}

func NewMsgRate(creator string, rideId string, ratee string, rating float32) *MsgRate {
	return &MsgRate{
		Creator: creator,
		RideId:  rideId,
		Ratee:   ratee,
		Rating:  rating,
	}
}

func (msg *MsgRate) Route() string {
	return RouterKey
}

func (msg *MsgRate) Type() string {
	return TypeMsgRate
}

func (msg *MsgRate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
