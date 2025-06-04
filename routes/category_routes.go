package routes

import (
	"store-app/controllers"
	"store-app/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(r *gin.Engine) {
	category := r.Group("/api/categories")
	category.Use(middleware.AdminOnlyMiddleware())
	{
		category.POST("/", controllers.CreateCategory)
		category.GET("/", controllers.GetAllCategories)
		category.GET("/:id", controllers.GetCategoryByID)
		category.PUT("/:id", controllers.UpdateCategory)
		category.DELETE("/:id", controllers.DeleteCategory)
	}
}
