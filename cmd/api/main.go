package main

import (
	"alura-go-base/app/auth"
	"alura-go-base/app/product"
	"alura-go-base/app/user"
	"alura-go-base/config"
	"alura-go-base/db"
	"alura-go-base/types"
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
	authRouter := auth.NewAuthRouter(authService, unSecureRouter)

	// =========== middleware ===========

	secureRouter.Use(func(h http.Handler) http.Handler {
		return auth.AuthMiddleware(userService, h)
	})

	server := NewServer(router, productRouter, userRouter, authRouter)

	server.ListenAndServe()
}
