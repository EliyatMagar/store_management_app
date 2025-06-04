package routes

import (
	"store-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}
}
