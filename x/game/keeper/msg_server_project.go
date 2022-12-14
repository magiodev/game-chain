package keeper

import (
	"context"
	"github.com/G4AL-Entertainment/g4al-chain/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"regexp"
	"strings"
)

func (k msgServer) CreateProject(goCtx context.Context, msg *types.MsgCreateProject) (*types.MsgCreateProjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := validateDeveloper(ctx, k, msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	err = validateArgsProject(msg.Symbol, msg.Description, msg.Name)
	if err != nil {
		return nil, err
	}

	// Regex first of all as we strip characters
	symbol := regExSymbol(msg.Symbol)

	// Check if the value already exists
	_, isFound := k.GetProject(
		ctx,
		symbol,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var project = types.Project{
		Creator:     msg.Creator,
		Symbol:      symbol,
		Name:        msg.Name,
		Description: msg.Description,
		Delegate:    msg.Delegate,
		CreatedAt:   int32(ctx.BlockHeight()),
		UpdatedAt:   int32(ctx.BlockHeight()),
	}

	k.SetProject(
		ctx,
		project,
	)
	return &types.MsgCreateProjectResponse{}, nil
}

func (k msgServer) UpdateProject(goCtx context.Context, msg *types.MsgUpdateProject) (*types.MsgUpdateProjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Skipping role policy validation as we use Ownership or Delegate array.string

	err := validateArgsProject(msg.Symbol, msg.Description, msg.Name)
	if err != nil {
		return nil, err
	}

	// Check if the value exists
	valFound, isFound := k.GetProject(
		ctx,
		regExSymbol(msg.Symbol), // regEx applied,  TODO consider that maybe here to update is not required, as we do not allow changing symbol
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Check if msg.Creator included in valFound.Delegate
	isDelegate := false
	for _, del := range valFound.Delegate {
		if del == msg.Creator {
			isDelegate = true
			break
		}
	}
	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator && !isDelegate {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner nor delegate address")
	}

	// TODO consider delegates removal also
	// appending new delegate values to valFound.Delegate
	for _, delegate := range msg.Delegate {
		bech32, err := sdk.AccAddressFromBech32(delegate)
		if err == nil { // TODO fix this, is not preventing invalid addresses to be appended
			valFound.Delegate = append(valFound.Delegate, bech32.String())
		}
	}

	valFound.Name = msg.Name
	valFound.Description = msg.Description
	valFound.UpdatedAt = int32(ctx.BlockHeight())

	k.SetProject(ctx, valFound)

	return &types.MsgUpdateProjectResponse{}, nil
}

// Private Methods

func validateDeveloper(ctx sdk.Context, k msgServer, msg *types.MsgCreateProject) error {
	// Checking administrator role
	val, found := k.permissionKeeper.GetDeveloper(ctx, msg.Creator)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "creator invalid developer address (%s)", msg.Creator)
	}
	if val.Blocked {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "creator developer address blocked (%s)", msg.Creator)
	}
	return nil
}

func regExSymbol(symbol string) string {
	reg, _ := regexp.Compile("([^\\w])")                       // leaving only words meaning azAZ09 and _ without spaces
	symbol = strings.ToLower(reg.ReplaceAllString(symbol, "")) // toLower
	return symbol
}

func validateArgsProject(symbol string, description string, name string) error {
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
