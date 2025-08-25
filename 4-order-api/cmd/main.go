package main

import (
	"go/courses/configs"
	"go/courses/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
}
