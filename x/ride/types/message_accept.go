package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAccept = "accept"

var _ sdk.Msg = &MsgAccept{}

func NewMsgAccept(creator string, idValue string) *MsgAccept {
	return &MsgAccept{
		Creator: creator,
		IdValue: idValue,
	}
}

func (msg *MsgAccept) Route() string {
	return RouterKey
}

func (msg *MsgAccept) Type() string {
	return TypeMsgAccept
}

func (msg *MsgAccept) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAccept) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAccept) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
