package main

import (
	"service_2/config"
	"service_2/internal/app"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %v\n", err)
	}
	fmt.Println(cfg)

	app.Run(cfg)
}
