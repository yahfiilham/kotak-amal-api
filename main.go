package main

import (
	"kotak-amal/config"
	"kotak-amal/server"

	"github.com/go-playground/validator/v10"
)

func main() {
	validation := validator.New()
	cfg := config.LoadConfig()
	dbInit, err := config.MySql(cfg)
	if err != nil {
		panic(err)
	}

	server := server.NewServer(dbInit, validation)
	server.ListenAndServe("3000")
}
