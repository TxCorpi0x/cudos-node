package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		MessageList:         []Message{},
		MessageSentList:     []uint64{},
		MessageReceivedList: []uint64{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in message
	messageIdMap := make(map[uint64]bool)
	messageCount := gs.GetMessageCount()
	for _, elem := range gs.MessageList {
		if _, ok := messageIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for message")
		}
		if elem.Id > messageCount {
			return fmt.Errorf("message id should be lower or equal than the last id")
		}
		messageIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
