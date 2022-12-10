package keeper

import (
	"context"
	"fmt"
	"github.com/G4AL-Entertainment/g4al-chain/x/denomfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"regexp"
	"strings"
)

func (k msgServer) CreateDenom(goCtx context.Context, msg *types.MsgCreateDenom) (*types.MsgCreateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := validateCreateDenom(ctx, k, msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Check if the value already exists
	_, isFound := k.GetDenom(
		ctx,
		msg.Symbol,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	if msg.Precision < 6 || msg.Precision > 18 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "precision must be a number between 6 and 18")
	}

	// Regex first of all as we strip characters
	symbol, err := regExSymbol(msg.Symbol)
	if err != nil {
		return nil, err
	}

	var denom = types.Denom{
		Creator:            msg.Creator,
		Symbol:             symbol,
		Project:            msg.Project,
		MaxSupply:          msg.MaxSupply,
		CanChangeMaxSupply: msg.CanChangeMaxSupply,
	}

	// Creating sdk.Coin
	//coin := sdk.NewCoin(denom.Symbol, math.NewInt(int64(0))) // improve this conversion
	// Minting coins with x/bank module
	//err = k.bankKeeper.MintCoins(ctx, "denomfactory", sdk.NewCoins(coin))
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "mint has not occurred")
	}
	// Creating metadata's dependencies
	var baseDenomUnit = banktypes.DenomUnit{
		Denom:    denom.Symbol,
		Exponent: 0,
	}
	var milliDenomUnit = banktypes.DenomUnit{
		Denom:    "m" + denom.Symbol,
		Exponent: uint32(msg.Precision) - 3, // TODO check
	}
	milliDenomUnit.Aliases = append(milliDenomUnit.Aliases, "milli"+denom.Symbol)
	var microDenomUnit = banktypes.DenomUnit{
		Denom:    "u" + denom.Symbol,
		Exponent: uint32(msg.Precision),
	}
	microDenomUnit.Aliases = append(microDenomUnit.Aliases, "micro"+denom.Symbol)
	// Creating bank.Metadata object
	var metadata = banktypes.Metadata{
		Description: msg.Description,
		Base:        denom.Symbol,
		Display:     msg.Name,
		Name:        msg.Name,
		Symbol:      denom.Symbol,
	}
	metadata.DenomUnits = append(metadata.DenomUnits, &baseDenomUnit)  // TODO check if & reference is correct
	metadata.DenomUnits = append(metadata.DenomUnits, &milliDenomUnit) // TODO check if & reference is correct
	metadata.DenomUnits = append(metadata.DenomUnits, &microDenomUnit) // TODO check if & reference is correct
	// Set metadata
	k.bankKeeper.SetDenomMetaData(ctx, metadata)

	k.SetDenom(
		ctx,
		denom,
	)
	return &types.MsgCreateDenomResponse{}, nil
}

func (k msgServer) UpdateDenom(goCtx context.Context, msg *types.MsgUpdateDenom) (*types.MsgUpdateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := validateUpdateDenom(ctx, k, msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Check if the value exists
	valFound, isFound := k.GetDenom(
		ctx,
		msg.Symbol,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// Allowing update maxSupply only if previously declared
	if valFound.CanChangeMaxSupply {
		valFound.MaxSupply = msg.MaxSupply
	}
	var denom = types.Denom{
		Creator:            valFound.Creator,
		Symbol:             valFound.Symbol,
		Project:            valFound.Project,
		MaxSupply:          valFound.MaxSupply,
		CanChangeMaxSupply: valFound.CanChangeMaxSupply,
	}

	// TODO implement metadata getting and setting here to update Name,Description, Uri, UriHash

	k.SetDenom(ctx, denom)

	return &types.MsgUpdateDenomResponse{}, nil
}

// Private Methods

func validateCreateDenom(ctx sdk.Context, k msgServer, msg *types.MsgCreateDenom) error {
	// Checking administrator role
	val, found := k.permissionKeeper.GetDeveloper(ctx, msg.Creator)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "creator invalid developer address (%s)", msg.Creator)
	}
	if val.Blocked {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "creator developer address blocked (%s)", msg.Creator)
	}

	// Checking project existing and related to this game developer or delegate
	project, found := k.gameKeeper.GetProject(ctx, msg.Project)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "project invalid symbol (%s)", msg.Project)
	}

	// Check if msg.Creator included in valFound.Delegate
	isDelegate := false
	for _, del := range project.Delegate {
		if del == msg.Creator {
			isDelegate = true
			break
		}
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != project.Creator && !isDelegate {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner nor delegate address")
	}
	return nil
}

func validateUpdateDenom(ctx sdk.Context, k msgServer, msg *types.MsgUpdateDenom) error {
	// Checking administrator role
	val, found := k.permissionKeeper.GetDeveloper(ctx, msg.Creator)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "creator invalid developer address (%s)", msg.Creator)
	}
	if val.Blocked {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "creator developer address blocked (%s)", msg.Creator)
	}

	// Checking project existing and related to this game developer or delegate
	project, found := k.gameKeeper.GetProject(ctx, msg.Project)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "project invalid symbol (%s)", msg.Project)
	}

	// Check if msg.Creator included in valFound.Delegate
	isDelegate := false
	for _, del := range project.Delegate {
		if del == msg.Creator {
			isDelegate = true
			break
		}
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != project.Creator && !isDelegate {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner nor delegate address")
	}
	return nil
}

func regExSymbol(symbol string) (string, error) {
	reg, _ := regexp.Compile("([^\\w])")                       // leaving only words meaning azAZ09 and _ without spaces
	symbol = strings.ToLower(reg.ReplaceAllString(symbol, "")) // toLower
	if len(symbol) < 3 {
		return "", sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("symbol must be at least 8 digits: (%s)", symbol)) // TODO remove "" as return value
	}
	return symbol, nil
}

// TODO mintDenom() to user
// TODO burnDenom() from user considering approval /authz module
// TODO	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
