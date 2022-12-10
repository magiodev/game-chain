package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyGenesisAdministrator = []byte("GenesisAdministrator")
	// TODO: Determine the default value
	DefaultGenesisAdministrator string = "genesis_administrator"
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	genesisAdministrator string,
) Params {
	return Params{
		GenesisAdministrator: genesisAdministrator,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultGenesisAdministrator,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyGenesisAdministrator, &p.GenesisAdministrator, validateGenesisAdministrator),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateGenesisAdministrator(p.GenesisAdministrator); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateGenesisAdministrator validates the GenesisAdministrator param
func validateGenesisAdministrator(v interface{}) error {
	genesisAdministrator, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = genesisAdministrator

	return nil
}
