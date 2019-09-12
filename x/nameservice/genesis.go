package nameservice

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	KeyValRecords []KeyVal `json:"keyVal_records"`
}

func NewGenesisState(keyValRecords []KeyVal) GenesisState {
	return GenesisState{KeyValRecords: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.KeyValRecords {
		if record.Owner == nil {
			return fmt.Errorf("invalid KeyValRecord: Value: %s. Error: Missing Owner", record.Value)
		}
		if record.Value == "" {
			return fmt.Errorf("invalid KeyValRecord: Owner: %s. Error: Missing Value", record.Owner)
		}
//		if record.Price == nil {
//			return fmt.Errorf("invalid WhoisRecord: Value: %s. Error: Missing Price", record.Value)
//		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		KeyValRecords: []KeyVal{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.KeyValRecords {
		keeper.SetKeyVal(ctx, record.Value, record)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []KeyVal
	iterator := k.GetKeysIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		key := string(iterator.Key())
		keyVal := k.GetKeyVal(ctx, key)
		records = append(records, keyVal)

	}
	return GenesisState{KeyValRecords: records}
}
