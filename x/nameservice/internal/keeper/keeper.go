package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/sdk-application-tutorial/x/nameservice/internal/types"
)

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
//	CoinKeeper bank.Keeper

	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the nameservice Keeper
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
//		CoinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

// Gets the entire Whois metadata struct for a name
func (k Keeper) GetKeyVal(ctx sdk.Context, key string) types.KeyVal {
	store := ctx.KVStore(k.storeKey)
	if !k.IsKeyPresent(ctx, key) {
		return types.NewKeyVal()
	}
	bz := store.Get([]byte(key))
	var keyVal types.KeyVal
	k.cdc.MustUnmarshalBinaryBare(bz, &keyVal)
	return keyVal
}

// Sets the entire Whois metadata struct for a name
func (k Keeper) SetKeyVal(ctx sdk.Context, key string, keyVal types.KeyVal) {
	if keyVal.Owner.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(key), k.cdc.MustMarshalBinaryBare(keyVal))
}

// Deletes the entire Whois metadata struct for a name
func (k Keeper) DeleteKeyVal(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(key))
}

// ResolveName - returns the string that the name resolves to
func (k Keeper) ResolveKey(ctx sdk.Context, key string) string {
	return k.GetKeyVal(ctx, key).Value
}

// SetName - sets the value string that a name resolves to
func (k Keeper) SetKey(ctx sdk.Context, key string, value string) {
	keyVal := k.GetKeyVal(ctx, key)
	keyVal.Value = value
	keyVal.Key = key
	k.SetKeyVal(ctx, key, keyVal)
}

// HasOwner - returns whether or not the name already has an owner
func (k Keeper) HasOwner(ctx sdk.Context, key string) bool {
	return !k.GetKeyVal(ctx, key).Owner.Empty()
}

// GetOwner - get the current owner of a name
func (k Keeper) GetOwner(ctx sdk.Context, key string) sdk.AccAddress {
	return k.GetKeyVal(ctx, key).Owner
}

// SetOwner - sets the current owner of a name
func (k Keeper) SetOwner(ctx sdk.Context, key string, owner sdk.AccAddress) {
	keyVal := k.GetKeyVal(ctx, key)
	keyVal.Owner = owner
	k.SetKeyVal(ctx, key, keyVal)
}

// GetPrice - gets the current price of a name
//func (k Keeper) GetPrice(ctx sdk.Context, name string) sdk.Coins {
//	return k.GetWhois(ctx, name).Price
//}

// SetPrice - sets the current price of a name
//func (k Keeper) SetPrice(ctx sdk.Context, name string, price sdk.Coins) {
//	whois := k.GetWhois(ctx, name)
//	whois.Price = price
//	k.SetWhois(ctx, name, whois)
//}

// Get an iterator over all names in which the keys are the names and the values are the whois
func (k Keeper) GetKeysIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

// Check if the name is present in the store or not
func (k Keeper) IsKeyPresent(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(key))
}
