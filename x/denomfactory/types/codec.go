package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateDenom{}, "denomfactory/CreateDenom", nil)
	cdc.RegisterConcrete(&MsgUpdateDenom{}, "denomfactory/UpdateDenom", nil)
	cdc.RegisterConcrete(&MsgMintDenom{}, "denomfactory/MintDenom", nil)
	cdc.RegisterConcrete(&MsgBurnDenom{}, "denomfactory/BurnDenom", nil)
	cdc.RegisterConcrete(&MsgTransferDenom{}, "denomfactory/TransferDenom", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDenom{},
		&MsgUpdateDenom{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMintDenom{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBurnDenom{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransferDenom{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
