package types

import "github.com/gorilla/mux"

type IHandler interface {
	RegisterHandlers(router *mux.Router)
}
