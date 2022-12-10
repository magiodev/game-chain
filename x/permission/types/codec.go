package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateAdministrator{}, "permission/CreateAdministrator", nil)
	cdc.RegisterConcrete(&MsgUpdateAdministrator{}, "permission/UpdateAdministrator", nil)
	cdc.RegisterConcrete(&MsgDeleteAdministrator{}, "permission/DeleteAdministrator", nil)
	cdc.RegisterConcrete(&MsgCreateDeveloper{}, "permission/CreateDeveloper", nil)
	cdc.RegisterConcrete(&MsgUpdateDeveloper{}, "permission/UpdateDeveloper", nil)
	cdc.RegisterConcrete(&MsgDeleteDeveloper{}, "permission/DeleteDeveloper", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateAdministrator{},
		&MsgUpdateAdministrator{},
		&MsgDeleteAdministrator{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDeveloper{},
		&MsgUpdateDeveloper{},
		&MsgDeleteDeveloper{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
