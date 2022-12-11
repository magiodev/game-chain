package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AdministratorList: []Administrator{},
		DeveloperList:     []Developer{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in administrator
	administratorIndexMap := make(map[string]struct{})

	for _, elem := range gs.AdministratorList {
		index := string(AdministratorKey(elem.Address))
		if _, ok := administratorIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for administrator")
		}
		administratorIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in developer
	developerIndexMap := make(map[string]struct{})

	for _, elem := range gs.DeveloperList {
		index := string(DeveloperKey(elem.Address))
		if _, ok := developerIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for developer")
		}
		developerIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
