package routes

import (
	"store-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterOrderItemRoutes(router *gin.Engine) {
	orderItem := router.Group("/api/order-items")
	{
		orderItem.POST("/", controllers.CreateOrderItem)
		orderItem.GET("/", controllers.GetOrderItems)
		orderItem.GET("/:id", controllers.GetOrderItemByID)
		orderItem.PUT("/:id", controllers.UpdateOrderItem)
		orderItem.DELETE("/:id", controllers.DeleteOrderItem)
	}
}
