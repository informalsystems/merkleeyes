package rest

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

const (
	restName = "key"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/keys", storeName), keysHandler(cliCtx, storeName)).Methods("GET")
//	r.HandleFunc(fmt.Sprintf("/%s/names", storeName), buyNameHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/keys", storeName), setKeyHandler(cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/keys/{%s}", storeName, restName), resolveKeyHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/keys/{%s}/keyVal", storeName, restName), keyValHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/keys", storeName), deleteKeyHandler(cliCtx)).Methods("DELETE")
}
