package routes

import (
	"store-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterInventoryLogRoutes(r *gin.Engine) {
	inventory := r.Group("/api/inventorylogs")
	{
		inventory.POST("/", controllers.CreateInventoryLog)
		inventory.GET("/", controllers.GetInventoryLogs)
		inventory.GET("/:id", controllers.GetInventoryLogByID)
		inventory.PUT("/:id", controllers.UpdateInventoryLog)
		inventory.DELETE("/:id", controllers.DeleteInventoryLog)
	}
}
