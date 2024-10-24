package types

import "github.com/gorilla/mux"

type Handler interface {
	RegisterHandlers(router *mux.Router)
}
