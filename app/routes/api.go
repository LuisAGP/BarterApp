package routes

import (
	"net/http"

	"github.com/LuisAGP/BarterApp/app/middlewares"
	"github.com/LuisAGP/BarterApp/app/services/users"
	"github.com/gin-gonic/gin"
)

func API() *gin.Engine {
	// Defining router
	r := gin.Default()

	// Public routes
	r.POST("/register", users.RegisterUser)
	r.POST("/login", users.LoginUser)

	// Protected routes
	authRoutes := r.Group("/api/v1")
	authRoutes.Use(middlewares.JWTAuthMiddleware())
	authRoutes.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong!"})
	})

	return r
}
