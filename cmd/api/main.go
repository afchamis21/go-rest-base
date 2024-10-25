package main

import (
	"alura-rest-base/app/auth"
	"alura-rest-base/app/product"
	"alura-rest-base/app/user"
	"alura-rest-base/config"
	"alura-rest-base/db"
	"alura-rest-base/types"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Port          int
	Host          string
	Router        *mux.Router
	ProductRouter types.IProductRouter
	UserRouter    types.IUserRouter
	AuthRouter    types.IAuthRouter
}

func (s Server) ListenAndServe() {
	log.Printf("Listening on %s:%d\n", s.Host, s.Port)
	http.ListenAndServe(fmt.Sprintf("%s:%d", s.Host, s.Port), s.Router)
}

func NewServer(
	router *mux.Router,
	productRouter types.IProductRouter,
	userRouter types.IUserRouter,
	authRouter types.IAuthRouter,
) *Server {
	return &Server{
		Port:          config.Envs.Port,
		Host:          config.Envs.PublicHost,
		Router:        router,
		ProductRouter: productRouter,
		UserRouter:    userRouter,
		AuthRouter:    authRouter,
	}
}

func main() {
	router := mux.NewRouter()
	secureRouter := router.PathPrefix("/api/v1/secure").Subrouter()
	unSecureRouter := router.PathPrefix("/api/v1/unsecure").Subrouter()

	database := db.ConnectToDatabase()

	// =========== repo, service, router ===========

	productStorage := product.NewProductRepo(database)
	productService := product.NewProductService(productStorage)
	productRouter := product.NewProductRouter(productService, secureRouter)

	userStorage := user.NewUserRepo(database)
	userService := user.NewUserService(userStorage)
	userRouter := user.NewUserRouter(userService, secureRouter)

	authService := auth.NewAuthService(userService)
	authRouter := auth.NewAuthRouter(unSecureRouter, authService, userService)

	// =========== middleware ===========

	secureRouter.Use(func(h http.Handler) http.Handler {
		return auth.AuthMiddleware(userService, h)
	})

	server := NewServer(router, productRouter, userRouter, authRouter)

	server.ListenAndServe()
}
