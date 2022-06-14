package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRequestRide{}, "ride/RequestRide", nil)
	cdc.RegisterConcrete(&MsgAccept{}, "ride/Accept", nil)
	cdc.RegisterConcrete(&MsgFinish{}, "ride/Finish", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRequestRide{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAccept{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFinish{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
