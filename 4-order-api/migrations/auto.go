package main

import (
	"go/courses/configs"
	"go/courses/internal/auth"
	"go/courses/internal/product"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	conf := configs.LoadConfig()

	db, err := gorm.Open(postgres.Open(conf.Db.Dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&product.Product{}, &auth.Auth{})
}
