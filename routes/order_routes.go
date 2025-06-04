package routes

import (
	"store-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(router *gin.Engine) {
	order := router.Group("/api/orders")
	{
		order.POST("/", controllers.CreateOrder)
		order.GET("/", controllers.GetOrders)
		order.GET("/:id", controllers.GetOrderByID)
		order.PUT("/:id", controllers.UpdateOrder)
		order.DELETE("/:id", controllers.DeleteOrder)
	}
}
