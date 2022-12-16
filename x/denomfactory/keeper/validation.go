package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) ValidateArgsDenom(symbol string, description string, name string) error {
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

func (k Keeper) ValidateProjectOwnershipOrDelegateByDenom(ctx sdk.Context, creator string, symbol string) error {
	// Checking if symbol/classID exists via x/nft
	denomFound, found := k.bankKeeper.GetDenomMetaData(ctx, symbol)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "symbol of denom not found x/nft (%s)", symbol)
	}

	// check on map to what project is associated with
	denomMapFound, found := k.GetDenom(ctx, denomFound.Symbol)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "symbol of denom not found x/assetfactory (%s)", symbol)
	}

	// Checking project existing and related to this game developer or delegate
	project, found := k.gameKeeper.GetProject(ctx, denomMapFound.Project)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "project invalid symbol (%s)", denomMapFound.Project)
	}

	// Check delegate
	err := k.gameKeeper.ValidateDelegate(creator, project)
	if err != nil {
		return err
	}
	return nil
}
