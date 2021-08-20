package main

import (
	"log"

	"github.com/majedutd990/bookstore-api/internal/config"
)

func init() {
	var cfg *config.Config
	if err := config.Parser("../../build/config/config.yaml", cfg); err != nil {
		log.Fatalln(err)
	}
	if err := config.ReadEnv(cfg); err != nil {
		log.Fatalln(err)
	}
	config.SetConfig(cfg)
}

func main() {

}
