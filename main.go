package main

import (
	"fmt"
	"os"

	"github.com/NagendraGokuwada/golang_jwt_auth/routes"
	routes "github.com/NagendraGokuwada/golang_jwt_auth/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, Nagendra!")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default port if not specified in the environment
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for the API-1"})
	})

	router.GET("/api2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for the API-2"})
	})
	router.Run(":" + port)
}
