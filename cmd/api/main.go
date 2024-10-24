package main

import (
	"alura-go-base/app/config"
	"alura-go-base/app/db"
	"alura-go-base/app/repo"
	"alura-go-base/app/services"
	"alura-go-base/server/middleware"
	"alura-go-base/server/routes"
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

	productStorage := repo.NewProductRepo(database)
	productService := services.NewProductService(productStorage)
	productRouter := routes.NewProductRouter(productService, secureRouter)

	userStorage := repo.NewUserRepo(database)
	userService := services.NewUserService(userStorage)
	userRouter := routes.NewUserRouter(userService, secureRouter)

	authService := services.NewAuthService(*userService)
	authRouter := routes.NewAuthRouter(authService, unSecureRouter)

	server := NewServer(router, productRouter, userRouter, authRouter)

	secureRouter.Use(func(h http.Handler) http.Handler {
		return middleware.AuthMiddleware(userService, h)
	})

	server.ListenAndServe()
}
