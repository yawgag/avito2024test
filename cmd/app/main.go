package main

import (
	"avito2024test/config"
	"avito2024test/internal/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	router := routes.SetupRoutes()

	fmt.Printf("starting server at: %s", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, router); err != nil {
		log.Fatal("Server start error: ", err)
	}
}
