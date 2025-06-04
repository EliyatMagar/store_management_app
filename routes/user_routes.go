package routes

import (
	"store-app/controllers"

	"store-app/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	users := r.Group("/api/users")
	users.Use(middleware.AdminOnlyMiddleware())
	{
		users.GET("/", controllers.GetAllUsers)
		users.GET("/:id", controllers.GetUserByID)
		users.PUT("/:id", controllers.UpdateUserByID)
		users.DELETE("/:id", controllers.DeleteUserByID)
	}
}
