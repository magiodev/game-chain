package utils

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/golang/protobuf/ptypes/wrappers"
	"regexp"
	"strings"
)

func RegExSymbol(symbol string) string {
	reg, _ := regexp.Compile("([^\\w])")                       // leaving only words meaning azAZ09 and _ without spaces
	symbol = strings.ToLower(reg.ReplaceAllString(symbol, "")) // toLower
	return symbol
}

func StringToAny(data string) (*codectypes.Any, error) {
	sv := &wrappers.StringValue{Value: data}
	msgData, err := codectypes.NewAnyWithValue(sv)
	if err != nil {
		return msgData, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "data not valid")
	}
	return msgData, nil
}
