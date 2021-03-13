package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/isnakolah/todoAPI/config"
	"github.com/isnakolah/todoAPI/migration"
	"github.com/isnakolah/todoAPI/route"
)

func init() {
	db := config.Init()
	migration.Migrate(db)
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := route.SetupRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
