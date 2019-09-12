package nameservice

import (
	"github.com/cosmos/sdk-application-tutorial/x/nameservice/internal/keeper"
	"github.com/cosmos/sdk-application-tutorial/x/nameservice/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper        = keeper.NewKeeper
	NewQuerier       = keeper.NewQuerier
//	NewMsgBuyName    = types.NewMsgBuyName
	NewMsgSetKey     = types.NewMsgSetKey
	NewMsgDeleteKey  = types.NewMsgDeleteKey
	NewKeyVal        = types.NewKeyVal
	ModuleCdc        = types.ModuleCdc
	RegisterCodec    = types.RegisterCodec
)

type (
	Keeper          = keeper.Keeper
	MsgSetKey       = types.MsgSetKey
//	MsgBuyName      = types.MsgBuyName
	MsgDeleteKey    = types.MsgDeleteKey
	QueryResResolve = types.QueryResResolve
	QueryResKeys    = types.QueryResKeys
	KeyVal          = types.KeyVal
)
