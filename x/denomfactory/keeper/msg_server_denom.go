package keeper

import (
	"context"
	"github.com/G4AL-Entertainment/g4al-chain/x/denomfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"regexp"
	"strings"
)

func (k msgServer) CreateDenom(goCtx context.Context, msg *types.MsgCreateDenom) (*types.MsgCreateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate
	err := validateDeveloper(ctx, k, msg.Creator)
	if err != nil {
		return nil, err
	}
	err = validateProjectOwnershipOrDelegateByProject(ctx, k, msg.Creator, msg.Symbol)
	if err != nil {
		return nil, err
	}

	// Check if input texts meet requirements
	err = validateArgsDenom(msg.Symbol, msg.Description, msg.Name, msg.Precision)
	if err != nil {
		return nil, err
	}

	// Check if the value already exists
	_, isFound := k.GetDenom(
		ctx,
		msg.Symbol,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	// Regex first of all as we strip characters
	symbol := regExSymbol(msg.Symbol)

	var denom = types.Denom{
		Creator:            msg.Creator,
		Symbol:             symbol,
		Project:            msg.Project,
		MaxSupply:          msg.MaxSupply,
		CanChangeMaxSupply: msg.CanChangeMaxSupply,
	}

	// Set metadata for Denom
	SetCoinMetadata(ctx, k, symbol, msg.Name, msg.Description, msg.Precision)

	k.SetDenom(
		ctx,
		denom,
	)
	return &types.MsgCreateDenomResponse{}, nil
}

func (k msgServer) UpdateDenom(goCtx context.Context, msg *types.MsgUpdateDenom) (*types.MsgUpdateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate
	err := validateDeveloper(ctx, k, msg.Creator)
	if err != nil {
		return nil, err
	}
	err = validateProjectOwnershipOrDelegateByProject(ctx, k, msg.Creator, msg.Symbol)
	if err != nil {
		return nil, err
	}

	// Check if the value exists
	valFound, isFound := k.GetDenom(
		ctx,
		msg.Symbol,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
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

	// Getting existing metadata
	var existingMetadata, found = k.bankKeeper.GetDenomMetaData(ctx, valFound.Symbol)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "metadata for denom not found")
	}
	//existingMetadata.Base // TODO use this as baseDenomString
	//existingMetadata.DenomUnits // TODO loop this to take proper value
	//TODO use properValue.Precision

	// Check if input texts meet requirements
	err = validateArgsDenom(msg.Symbol, msg.Description, msg.Name, 6) // TODO de-hardcode this
	if err != nil {
		return nil, err
	}

	// Set metadata for Denom
	SetCoinMetadata(ctx, k, existingMetadata.Symbol, msg.Name, msg.Description, 6) // TODO de-hardcode this

	k.SetDenom(ctx, denom)

	return &types.MsgUpdateDenomResponse{}, nil
}

// Private Methods

func SetCoinMetadata(ctx sdk.Context, k msgServer, symbol string, name string, description string, precision int32) {
	// Creating metadata TODO check exponent assign, consider limiting only to fixed number of decimals for all the denoms created
	var baseDenomUnit = banktypes.DenomUnit{
		Denom:    symbol,
		Exponent: 0,
	}
	var milliDenomUnit = banktypes.DenomUnit{
		Denom:    "m" + symbol,
		Exponent: uint32(precision) - 3, // TODO check
	}
	milliDenomUnit.Aliases = append(milliDenomUnit.Aliases, "milli"+symbol)
	var microDenomUnit = banktypes.DenomUnit{
		Denom:    "u" + symbol,
		Exponent: uint32(precision),
	}
	microDenomUnit.Aliases = append(microDenomUnit.Aliases, "micro"+symbol)
	// Creating bank.Metadata object
	var metadata = banktypes.Metadata{
		Description: description,
		Base:        symbol,
		Display:     name,
		Name:        name,
		Symbol:      symbol,
	}
	metadata.DenomUnits = append(metadata.DenomUnits, &baseDenomUnit)  // TODO check if & reference is correct
	metadata.DenomUnits = append(metadata.DenomUnits, &milliDenomUnit) // TODO check if & reference is correct
	metadata.DenomUnits = append(metadata.DenomUnits, &microDenomUnit) // TODO check if & reference is correct
	// Set metadata
	k.bankKeeper.SetDenomMetaData(ctx, metadata)
}

// Private Methods

func validateDeveloper(ctx sdk.Context, k msgServer, creator string) error {
	// Checking developer role
	val, found := k.permissionKeeper.GetDeveloper(ctx, creator)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "creator invalid developer address (%s)", creator)
	}
	if val.Blocked {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "creator developer address blocked (%s)", creator)
	}
	return nil
}

func validateProjectOwnershipOrDelegateByProject(ctx sdk.Context, k msgServer, creator string, symbol string) error {
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

func validateArgsDenom(symbol string, description string, name string, precision int32) error {
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
	if precision < PrecisionMinValue {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "precision is needed at least %d value", PrecisionMinValue)
	}
	if precision > PrecisionMaxValue {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "precision is needed at most %d value", PrecisionMaxValue)
	}
	return nil
}

func regExSymbol(symbol string) string {
	reg, _ := regexp.Compile("([^\\w])")                       // leaving only words meaning azAZ09 and _ without spaces
	symbol = strings.ToLower(reg.ReplaceAllString(symbol, "")) // toLower
	return symbol
}

// TODO mintDenom() to user
// TODO burnDenom() from user considering approval /authz module
// TODO	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
