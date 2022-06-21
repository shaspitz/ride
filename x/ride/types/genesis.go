package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		NextRide: &NextRide{
			IdValue:  uint64(DefaultIndex),
			FifoHead: NoFifoIdKey,
			FifoTail: NoFifoIdKey,
		},
		StoredRideList:   []StoredRide{},
		RatingStructList: []RatingStruct{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in storedRide
	storedRideIndexMap := make(map[string]struct{})

	for _, elem := range gs.StoredRideList {
		index := string(StoredRideKey(elem.Index))
		if _, ok := storedRideIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for storedRide")
		}
		storedRideIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in ratingStruct
	ratingStructIndexMap := make(map[string]struct{})

	for _, elem := range gs.RatingStructList {
		index := string(RatingStructKey(elem.Index))
		if _, ok := ratingStructIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for ratingStruct")
		}
		ratingStructIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
