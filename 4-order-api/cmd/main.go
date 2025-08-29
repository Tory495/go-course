package main

import (
	"fmt"
	"go/courses/configs"
	"go/courses/internal/auth"
	"go/courses/internal/product"
	"go/courses/pkg/db"
	"go/courses/pkg/middleware"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	conf := configs.LoadConfig()
	db := db.NewDb(conf)

	router := http.NewServeMux()

	// Repositories

	authRepository := auth.NewAuthRepository(db)

	productRepository := product.NewProductRepository(product.ProductRepositoryDeps{
		Database: db,
	})

	// Services

	authService := auth.NewAuthService(authRepository)

	// Handlers

	auth.NewAuthHandler(router, &auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})

	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepository,
	})

	// Middlewares
	stack := middleware.Chain(middleware.Logging)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("Server is listening to port 8081")
	server.ListenAndServe()
}
