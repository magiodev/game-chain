package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) ValidateAdministrator(ctx sdk.Context, creator string) error {
	// Checking administrator role
	val, found := k.GetAdministrator(ctx, creator)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "creator invalid administrator address (%s)", creator)
	}
	if val.Blocked {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "creator administrator address blocked (%s)", creator)
	}
	return nil
}

func (k Keeper) ValidateUpdateAdministrator(ctx sdk.Context, creator string, address string) error {
	// Checking if last admin
	if len(k.GetAllAdministrator(ctx)) < 2 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "deletion of the last administrator is not permitted")
	}
	// Check if the value exists
	toDelete, isFound := k.GetAdministrator(ctx, address)
	if !isFound {
		return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "to delete index not set")
	}
	// Checks if the msg creator is the same as the toDelete admin owner
	if creator != toDelete.Address {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "an administrator cannot delete himself")
	}
	return nil
}
