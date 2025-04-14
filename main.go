package main

import (
	"net/http"

	"github.com/LuisAGP/BarterApp/db"
	"github.com/LuisAGP/BarterApp/handlers"
	"github.com/LuisAGP/BarterApp/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {

	// Apply migrations
	db.AutoMigrate()

	// Define router
	r := gin.Default()

	// Public routes
	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.LoginUser)

	// Protected routes
	authRoutes := r.Group("/api/v1")
	authRoutes.Use(middlewares.JWTAuthMiddleware())
	authRoutes.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong!"})
	})

	r.Run(":8080")

}
