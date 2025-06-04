package routes

import (
	"store-app/controllers"
	"store-app/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterProductRoutes registers all product-related routes
func RegisterProductRoutes(router *gin.Engine) {
	productGroup := router.Group("/api/products")
	productGroup.Use(middleware.AdminOnlyMiddleware()) // Apply authentication middleware
	{
		productGroup.POST("/", controllers.CreateProduct)
		productGroup.POST("/bulk", controllers.CreateMultipleProducts) // Added bulk create route
		productGroup.GET("/", controllers.GetAllProducts)
		productGroup.GET("/:id", controllers.GetProductByID)
		productGroup.PUT("/:id", controllers.UpdateProduct)
		productGroup.DELETE("/:id", controllers.DeleteProduct)

		// Additional functionality
		productGroup.GET("/search", controllers.SearchProducts)
		productGroup.GET("/category/:category_id", controllers.GetProductsByCategory)
		productGroup.GET("/:id/stock", controllers.GetProductStock)
	}
}
