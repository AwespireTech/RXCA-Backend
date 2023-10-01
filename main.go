package main

import (
	"os"

	"github.com/AwespireTech/dXCA-Backend/database"
	"github.com/AwespireTech/dXCA-Backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database_url := os.Getenv("DATABASE_URL")
	database.Init(database_url)
	router := createRouter()
	router.Run(":8080")
}

func createRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to dXCA API",
		})
	})
	api := router.Group("/api")
	routes.SetDAORoutes(api)

	return router
}