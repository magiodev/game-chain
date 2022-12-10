package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateClass{}, "assetfactory/CreateClass", nil)
	cdc.RegisterConcrete(&MsgUpdateClass{}, "assetfactory/UpdateClass", nil)
	cdc.RegisterConcrete(&MsgMintNft{}, "assetfactory/MintNft", nil)
	cdc.RegisterConcrete(&MsgUpdateNft{}, "assetfactory/UpdateNft", nil)
	cdc.RegisterConcrete(&MsgBurnNft{}, "assetfactory/BurnNft", nil)
	cdc.RegisterConcrete(&MsgTransferNft{}, "assetfactory/TransferNft", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateClass{},
		&MsgUpdateClass{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMintNft{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateNft{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBurnNft{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransferNft{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
