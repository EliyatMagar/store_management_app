package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"store-app/database"
	"store-app/models"
)

//Create Order Item

func CreateOrderItem(c *gin.Context) {
	var orderItem models.OrderItem

	if err := c.ShouldBindJSON(&orderItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := database.DB.Create(&orderItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
	}
	c.JSON(http.StatusOK, orderItem)
}

//Get All order Items

func GetOrderItems(c *gin.Context) {
	var orderItems []models.OrderItem

	if err := database.DB.Preload("Order").Preload("Product").Find(&orderItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order items"})
		return
	}
	c.JSON(http.StatusOK, orderItems)
}

//Get Order Item by ID

func GetOrderItemByID(c *gin.Context) {
	id := c.Param("id")
	var orderItem models.OrderItem

	if err := database.DB.Preload("Order").Preload("Product").First(&orderItem, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order item not found"})
	}
}

//Update Order Item

func UpdateOrderItem(c *gin.Context) {
	id := c.Param("id")

	var orderItem models.OrderItem

	if err := database.DB.First(&orderItem, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order item not found"})
		return
	}
}

//Delete Order Item

func DeleteOrderItem(c *gin.Context) {
	id := c.Param("id")
	var orderItem models.OrderItem

	if err := database.DB.First(&orderItem, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order item not found"})
		return
	}

	if err := database.DB.Delete(&orderItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order item deleted successfully"})
}
