package keeper

import (
	"context"
	"fmt"
	"github.com/G4AL-Entertainment/g4al-chain/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"regexp"
	"strings"
)

func (k msgServer) CreateProject(goCtx context.Context, msg *types.MsgCreateProject) (*types.MsgCreateProjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := validateCreateProject(ctx, k, msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Regex first of all as we strip characters
	symbol, err := regExSymbol(msg.Symbol)
	if err != nil {
		return nil, err
	}

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

	// Regex first of all as we strip characters
	symbol, err := regExSymbol(msg.Symbol)
	if err != nil {
		return nil, err
	}

	// Check if the value exists
	valFound, isFound := k.GetProject(
		ctx,
		symbol,
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

	// appending new delegate values to valFound.Delegate
	for _, delegate := range msg.Delegate {
		bech32, err := sdk.AccAddressFromBech32(delegate)
		if err == nil {
			valFound.Delegate = append(valFound.Delegate, bech32.String())
		}
	}

	var project = types.Project{
		Creator:     valFound.Creator,
		Symbol:      valFound.Symbol,
		Name:        msg.Name,
		Description: msg.Description,
		Delegate:    valFound.Delegate, // updated values
		CreatedAt:   valFound.CreatedAt,
		UpdatedAt:   int32(ctx.BlockHeight()),
	}

	k.SetProject(ctx, project)

	return &types.MsgUpdateProjectResponse{}, nil
}

// Private Methods

func validateCreateProject(ctx sdk.Context, k msgServer, msg *types.MsgCreateProject) error {
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

func regExSymbol(symbol string) (string, error) {
	reg, _ := regexp.Compile("([^\\w])")                       // leaving only words meaning azAZ09 and _ without spaces
	symbol = strings.ToLower(reg.ReplaceAllString(symbol, "")) // toLower
	if len(symbol) < 8 {
		return "", sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("symbol must be at least 8 digits: (%s)", symbol)) // TODO remove "" as return value
	}
	return symbol, nil
}
