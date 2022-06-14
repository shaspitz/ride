package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgFinish = "finish"

var _ sdk.Msg = &MsgFinish{}

func NewMsgFinish(creator string, idValue string, location string) *MsgFinish {
	return &MsgFinish{
		Creator:  creator,
		IdValue:  idValue,
		Location: location,
	}
}

func (msg *MsgFinish) Route() string {
	return RouterKey
}

func (msg *MsgFinish) Type() string {
	return TypeMsgFinish
}

func (msg *MsgFinish) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFinish) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFinish) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
