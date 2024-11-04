package main

import (
	"example/restapi/internal/api"
	"example/restapi/internal/config"
	"example/restapi/internal/database"
	"log"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Database Connection error: %v", err)
	}
	d, _ := db.DB()
	defer d.Close()

	r := api.SetUpRouter(db)

	log.Println("Server running on", cfg.ServerAddress)
	if err := r.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
