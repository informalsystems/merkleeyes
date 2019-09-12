package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeKeyDoesNotExist sdk.CodeType = 101
)

func ErrKeyDoesNotExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeKeyDoesNotExist, "Key does not exist")
}
