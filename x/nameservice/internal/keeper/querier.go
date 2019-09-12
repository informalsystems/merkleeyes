package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/sdk-application-tutorial/x/nameservice/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the nameservice Querier
const (
	QueryResolve = "resolve"
	QueryKeyVal   = "keyVal"
	QueryKeys   = "keys"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryResolve:
			return queryResolve(ctx, path[1:], req, keeper)
		case QueryKeyVal:
			return queryKeyVal(ctx, path[1:], req, keeper)
		case QueryKeys:
			return queryKeys(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown nameservice query endpoint")
		}
	}
}

// nolint: unparam
func queryResolve(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	value := keeper.ResolveKey(ctx, path[0])

	if value == "" {
		return []byte{}, sdk.ErrUnknownRequest("could not resolve key")
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResResolve{Value: value})
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

// nolint: unparam
func queryKeyVal(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	keyVal := keeper.GetKeyVal(ctx, path[0])

	res, err := codec.MarshalJSONIndent(keeper.cdc, keyVal)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryKeys(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var keysList types.QueryResKeys

	iterator := keeper.GetKeysIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		keysList = append(keysList, string(iterator.Key()))
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, keysList)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
