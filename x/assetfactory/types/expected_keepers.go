package types

import (
	gametypes "github.com/G4AL-Entertainment/g4al-chain/x/game/types"
	permissiontypes "github.com/G4AL-Entertainment/g4al-chain/x/permission/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
)

type PermissionKeeper interface {
	// Methods imported from permission should be defined here
	GetDeveloper(ctx sdk.Context, address string) (val permissiontypes.Developer, found bool)
	ValidateDeveloper(ctx sdk.Context, address string) error
}

type GameKeeper interface {
	// Methods imported from game should be defined here
	GetProject(ctx sdk.Context, symbol string) (val gametypes.Project, found bool)
	ValidateDelegate(creator string, project gametypes.Project) error
	ValidateProjectOwnershipOrDelegateByProject(ctx sdk.Context, creator string, symbol string) error
}

type NftKeeper interface {
	// Class
	SaveClass(ctx sdk.Context, class nft.Class) error
	UpdateClass(ctx sdk.Context, class nft.Class) error

	// NFT
	Mint(ctx sdk.Context, nft nft.NFT, receiver sdk.AccAddress) error // updates totalSupply
	Burn(ctx sdk.Context, classId string, nftId string) error         // updates totalSupply
	Update(ctx sdk.Context, nft nft.NFT) error
	Transfer(ctx sdk.Context, classId string, nftId string, receiver sdk.AccAddress) error

	// Getters
	GetClass(ctx sdk.Context, classId string) (nft.Class, bool)
	GetClasses(ctx sdk.Context) (classes []*nft.Class)
	GetNFT(ctx sdk.Context, classID string, nftID string) (nft.NFT, bool)
	GetNFTsOfClassByOwner(ctx sdk.Context, classId string, owner sdk.AccAddress) []nft.NFT
	GetNFTsOfClass(ctx sdk.Context, classId string) []nft.NFT
	GetOwner(ctx sdk.Context, classId string, nftId string) sdk.AccAddress
	GetBalance(ctx sdk.Context, classId string, owner sdk.AccAddress) uint64
	GetTotalSupply(ctx sdk.Context, classId string) uint64
	// Methods imported from nft should be defined here

	nft.QueryServer
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}
