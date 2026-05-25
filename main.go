package main

import (
	"log"

	"hello/config"
	"hello/routes"
)

func main() {
	if err := config.ConnectDatabase(); err != nil {
		log.Fatal("Database setup failed:", err)
	}

	router := routes.SetupRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
