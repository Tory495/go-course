package main

import (
	"fmt"
	"go/courses/configs"
	"go/courses/internal/product"
	"go/courses/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)

	router := http.NewServeMux()

	productRepository := product.NewProductRepository(product.ProductRepositoryDeps{
		Database: db,
	})

	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening to port 8081")
	server.ListenAndServe()
}
