package types

import (
	gametypes "github.com/G4AL-Entertainment/g4al-chain/x/game/types"
	permissiontypes "github.com/G4AL-Entertainment/g4al-chain/x/permission/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type PermissionKeeper interface {
	GetDeveloper(ctx sdk.Context, address string) (val permissiontypes.Developer, found bool)
	ValidateDeveloper(ctx sdk.Context, address string) error
}

type GameKeeper interface {
	GetProject(ctx sdk.Context, symbol string) (val gametypes.Project, found bool)
	ValidateProjectOwnershipOrDelegateByProject(ctx sdk.Context, creator string, symbol string) error
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	// TODO remove methods restricting to only essential ones, security flaws
	keeper.SendKeeper
	WithMintCoinsRestriction(keeper.MintingRestrictionFn) keeper.BaseKeeper

	InitGenesis(sdk.Context, *banktypes.GenesisState)
	ExportGenesis(sdk.Context) *banktypes.GenesisState

	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	HasSupply(ctx sdk.Context, denom string) bool
	GetPaginatedTotalSupply(ctx sdk.Context, pagination *query.PageRequest) (sdk.Coins, *query.PageResponse, error)
	IterateTotalSupply(ctx sdk.Context, cb func(sdk.Coin) bool)
	GetDenomMetaData(ctx sdk.Context, denom string) (banktypes.Metadata, bool)
	HasDenomMetaData(ctx sdk.Context, denom string) bool
	SetDenomMetaData(ctx sdk.Context, denomMetaData banktypes.Metadata)
	IterateAllDenomMetaData(ctx sdk.Context, cb func(banktypes.Metadata) bool)

	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	DelegateCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	UndelegateCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error

	DelegateCoins(ctx sdk.Context, delegatorAddr, moduleAccAddr sdk.AccAddress, amt sdk.Coins) error
	UndelegateCoins(ctx sdk.Context, moduleAccAddr, delegatorAddr sdk.AccAddress, amt sdk.Coins) error // Methods imported from bank should be defined here
}
