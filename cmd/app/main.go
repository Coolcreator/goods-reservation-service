package main

import (
	"log"

	"github.com/1nkh3art1/goods-reservation-service/config"
	"github.com/1nkh3art1/goods-reservation-service/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
