package main

import (
	"log"

	"github.com/majedutd990/bookstore-api/internal/application"
	"github.com/majedutd990/bookstore-api/internal/config"
)

var cfg = &config.Config{}

func init() {
	if err := config.Parser("../../build/config/config.yaml", cfg); err != nil {
		log.Fatalln(err)
	}
	if err := config.ReadEnv(cfg); err != nil {
		log.Fatalln(err)
	}
	config.SetConfig(cfg)
}

func main() {
	if err := application.Run(cfg); err != nil {
		log.Fatalln(err)
	}
}
