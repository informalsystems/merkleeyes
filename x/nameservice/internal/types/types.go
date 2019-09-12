package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)


// whois is a struct that contains all the metadata of a name
type KeyVal struct {
	Key string	     `json:"key"`
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`

}

// Returns a new Whois with the minprice as the price
func NewKeyVal() KeyVal {
	return KeyVal{
		Key: "",
		Value: "",
	}
}

// implement fmt.Stringer
func (w KeyVal) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Value: %s Key: %s`, w.Owner, w.Value, w.Key))
}
