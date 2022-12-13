package keeper

import (
	"context"

	"github.com/golang/protobuf/ptypes/wrappers"

	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/nft"
)

type Constants struct {
	MinDescriptionLength int
	MaxDescriptionLength int
	MinNameLength        int
	MaxNameLength        int
	MinSymbolLength      int
	MaxSymbolLength      int
}

var Consts Constants = Constants{
	MinDescriptionLength: 0,
	MaxDescriptionLength: 0,
	MinNameLength:        0,
	MaxNameLength:        0,
	MinSymbolLength:      0,
	MaxSymbolLength:      0,
}

func (k msgServer) CreateClass(goCtx context.Context, msg *types.MsgCreateClass) (*types.MsgCreateClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks if the msg creator is a developer
	_, isDevFound := k.permissionKeeper.GetDeveloper(ctx, msg.Creator)
	if !isDevFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "caller is not a developer")
	}

	// Checks if game exists
	game, isGameFound := k.gameKeeper.GetProject(ctx, msg.Project)
	if !isGameFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "game does not exist")
	}

	// Checks if masg creator is the game creator
	if game.Creator != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "game does not exist")
	}

	// Parameter validation

	// Check if the value already exists
	_, isFound := k.GetClass(
		ctx,
		msg.Symbol,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	// Check if input texts meet requirements
	err := ValidateInputTextClass(msg.Symbol, msg.Description, msg.Name)
	if err != nil {
		return nil, err
	}

	var class = types.Class{
		Creator:            msg.Creator,
		Symbol:             msg.Symbol,
		Project:            msg.Project,
		MaxSupply:          msg.MaxSupply,
		CanChangeMaxSupply: msg.CanChangeMaxSupply,
	}

	// Treating msg.Data any value
	msgData, err := StringToAny(msg.Data)
	if err != nil {
		return nil, err
	}

	var nftClass = nft.Class{
		Id:          msg.Symbol,
		Name:        msg.Name,
		Symbol:      msg.Symbol,
		Description: msg.Description,
		Uri:         msg.Uri,
		UriHash:     msg.UriHash,
		Data:        msgData,
	}
	err = k.nftKeeper.SaveClass(ctx, nftClass)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "class creation has not occurred")
	}

	k.SetClass(
		ctx,
		class,
	)
	return &types.MsgCreateClassResponse{}, nil
}

func (k msgServer) UpdateClass(goCtx context.Context, msg *types.MsgUpdateClass) (*types.MsgUpdateClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetClass(
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

	err := ValidateInputTextClass(msg.Symbol, msg.Description, msg.Name)
	if err != nil {
		return nil, err
	}

	var class = types.Class{
		Creator:   msg.Creator,
		Symbol:    msg.Symbol,
		MaxSupply: msg.MaxSupply,
	}

	//TODO metadata workflow

	k.SetClass(ctx, class)

	return &types.MsgUpdateClassResponse{}, nil
}

// StringToAny TODO check if working properly
func StringToAny(data string) (*codectypes.Any, error) {
	sv := &wrappers.StringValue{Value: data}
	msgData, err := codectypes.NewAnyWithValue(sv)
	if err != nil {
		return msgData, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "data not valid")
	}
	return msgData, nil
}

func ValidateInputTextClass(symbol string, description string, name string) error {
	if len(symbol) < Consts.MinSymbolLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "symbol is needed and must contain at least %d characters", Consts.MinSymbolLength)
	}
	if len(symbol) > Consts.MaxSymbolLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "symbol is needed and can contain at most %d characters", Consts.MaxSymbolLength)
	}

	if len(description) < Consts.MinDescriptionLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "description is needed and must contain at least %d characters", Consts.MinDescriptionLength)
	}
	if len(description) > Consts.MaxDescriptionLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "description is needed and must contain at most %d characters", Consts.MaxDescriptionLength)
	}

	if len(name) < Consts.MinNameLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "name is needed and must contain at least %d characters", Consts.MinNameLength)
	}
	if len(name) > Consts.MaxNameLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "name is needed and can contain at most %d characters", Consts.MaxNameLength)
	}
	return nil
}
