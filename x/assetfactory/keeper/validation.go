package keeper

import (
	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) ValidateMaxSupply(ctx sdk.Context, symbol string) error {
	// check on map to what project is associated with TODO this is repeated, remove
	class, found := k.GetClass(ctx, symbol)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "symbol of classID not found x/assetfactory (%s)", symbol)
	}
	// Validate maxSupply only if not 0
	if class.MaxSupply != 0 {
		// If the current total supply is already equal (or higher just in case, but should do not happen)
		if k.nftKeeper.GetTotalSupply(ctx, class.Symbol) >= uint64(class.MaxSupply) {
			return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "max supply already reached for class (%s)", symbol)
		}
	}
	return nil
}

func (k Keeper) ValidateNftOwnershipOrAllowance(ctx sdk.Context, msg *types.MsgTransferNft) error {
	// Check if nft owner is msg owner
	owner := k.nftKeeper.GetOwner(ctx, msg.Symbol, msg.Id)
	if owner.String() != msg.Creator {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "nft owner is not msg creator")
	}

	// TODO Validate allowance (see x/authz if can help with nfts?) (or maybe skip ownership if developer. but not delegate?)

	return nil
}

func (k Keeper) ValidateProjectOwnershipOrDelegateByClassId(ctx sdk.Context, creator string, symbol string) error {
	// Checking if symbol/classID exists via x/nft
	classFound, found := k.nftKeeper.GetClass(ctx, symbol)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "symbol of classID not found x/nft (%s)", symbol)
	}

	// check on map to what project is associated with
	classMapFound, found := k.GetClass(ctx, classFound.Symbol)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "symbol of classID not found x/assetfactory (%s)", symbol)
	}

	// Checking project existing and related to this game developer or delegate
	project, found := k.gameKeeper.GetProject(ctx, classMapFound.Project)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "project invalid symbol (%s)", classMapFound.Project)
	}

	// Check if msg.Creator included in valFound.Delegate
	isDelegate := false
	for _, del := range project.Delegate {
		if del == creator {
			isDelegate = true
			break
		}
	}
	// Checks if the msg creator is the same as the current owner
	if creator != project.Creator && !isDelegate {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner nor delegate address")
	}
	return nil
}

func (k Keeper) ValidateArgsClass(symbol string, description string, name string) error {
	if len(symbol) < SymbolMinLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "symbol is needed and must contain at least %d characters", SymbolMinLength)
	}
	if len(symbol) > SymbolMaxLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "symbol is needed and can contain at most %d characters", SymbolMaxLength)
	}
	if len(name) < NameMinLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "name is needed and must contain at least %d characters", NameMinLength)
	}
	if len(name) > NameMaxLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "name is needed and can contain at most %d characters", NameMaxLength)
	}
	if len(description) < DescriptionMinLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "description is needed and must contain at least %d characters", DescriptionMinLength)
	}
	if len(description) > DescriptionMaxLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "description is needed and must contain at most %d characters", DescriptionMaxLength)
	}
	return nil
}

func (k Keeper) ValidateProjectOwnershipOrDelegateByProject(ctx sdk.Context, creator string, symbol string) error {
	// Checking project existing and related to this game developer or delegate
	project, found := k.gameKeeper.GetProject(ctx, symbol)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "project invalid symbol (%s)", symbol)
	}

	// Check if msg.Creator included in valFound.Delegate
	isDelegate := false
	for _, del := range project.Delegate {
		if del == creator {
			isDelegate = true
			break
		}
	}
	// Checks if the msg creator is the same as the current owner
	if creator != project.Creator && !isDelegate {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner nor delegate address")
	}
	return nil
}
