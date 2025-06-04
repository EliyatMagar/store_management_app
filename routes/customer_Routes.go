package routes

import (
	"store-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterCustomerRoutes(r *gin.Engine) {
	customer := r.Group("/api/customers")
	{
		customer.GET("/", controllers.GetCustomers)
		customer.GET("/:id", controllers.GetCustomerByID)
		customer.POST("/", controllers.CreateCustomer)
		customer.POST("/bulk", controllers.CreateMultipleCustomers)

		customer.PUT("/:id", controllers.UpdateCustomer)
		customer.DELETE("/:id", controllers.DeleteCustomer)
	}
}
