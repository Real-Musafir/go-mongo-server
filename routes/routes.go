package routes

import (
	"github.com/Real-Musafir/bookshop/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine){
	router := r.Group("/")

	// Register error handler middleware
	router.Use(middleware.ErrorHandler())

	userRoutes := router.Group("/users")

	{
		RegisterUserRoutes(userRoutes)
	}

	authRoutes := router.Group("/auth")
	{
		RegisterAuthRoutes(authRoutes)
	}
}