package main

import (
	"example/restapi/internal/config"
	"example/restapi/internal/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Database Connection error: ", err)
	}
	d, _ := db.DB()
	defer d.Close()
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "Hello",
		})
	})

	log.Println("Server running on", cfg.ServerAddress)
	if err := r.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
