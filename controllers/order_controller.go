package controllers

import (
	"net/http"
	"time"

	"store-app/database"
	"store-app/models"

	"github.com/gin-gonic/gin"
)

//Create Order

func CreateOrder(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	order.OrderDate = time.Now() //Set order date

	if err := database.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create order"})
		return
	}

	c.JSON(http.StatusOK, order)
}

//Get All Orders

func GetOrders(c *gin.Context) {
	var orders []models.Order

	if err := database.DB.Preload("Customer").Preload("User").Preload("OrderItems").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

//Get Single Order

func GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	if err := database.DB.Preload("Customer").Preload("User").Preload("OrderItems").First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

//Update Order

func UpdateOrder(c *gin.Context) {
	id := c.Param("id")

	var order models.Order

	if err := database.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	//Update allowed fields only
	order.Status = input.Status
	order.TotalAmount = input.TotalAmount
	order.ShippingAddress = input.ShippingAddress
	order.PaymentMethod = input.PaymentMethod

	if err := database.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	c.JSON(http.StatusOK, order)
}

//Delete Order

func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	if err := database.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if err := database.DB.Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
