package controllers

import (
	"net/http"
	"store-app/database"
	"store-app/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateInventoryLog(c *gin.Context) {
	var input models.InventoryLogInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.QuantityChange == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity change must be non-zero"})
		return
	}

	validTypes := map[string]bool{"addition": true, "removal": true, "adjustment": true}
	if !validTypes[input.Type] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type"})
		return
	}

	var product models.Product
	if err := database.DB.Where("name = ?", input.ProductName).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}

	var user models.User
	if err := database.DB.Where("id = ?", input.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserID not found"})
		return
	}

	log := models.InventoryLog{
		ProductID:      product.ID,
		UserID:         user.ID,
		QuantityChange: input.QuantityChange,
		Type:           input.Type,
		Date:           time.Now(),
	}

	if err := database.DB.Create(&log).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create inventory log"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Inventory log created successfully", "data": log})
}

func GetInventoryLogs(c *gin.Context) {
	var logs []models.InventoryLog
	if err := database.DB.Preload("Product").Preload("User").Find(&logs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve inventory logs"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": logs})
}

func GetInventoryLogByID(c *gin.Context) {
	id := c.Param("id")
	var log models.InventoryLog
	if err := database.DB.Preload("Product").Preload("User").First(&log, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory log not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": log})
}

func UpdateInventoryLog(c *gin.Context) {
	id := c.Param("id")
	var input models.InventoryLogInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.QuantityChange == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity change must be non-zero"})
		return
	}

	validTypes := map[string]bool{"addition": true, "removal": true, "adjustment": true}
	if !validTypes[input.Type] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type"})
		return
	}

	var log models.InventoryLog
	if err := database.DB.First(&log, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory log not found"})
		return
	}

	var product models.Product
	if err := database.DB.Where("name = ?", input.ProductName).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}

	var user models.User
	if err := database.DB.Where("id = ?", input.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	log.ProductID = product.ID
	log.UserID = user.ID
	log.QuantityChange = input.QuantityChange
	log.Type = input.Type
	log.Date = time.Now() // or preserve original date if you want

	if err := database.DB.Save(&log).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory log"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory log updated successfully", "data": log})
}

func DeleteInventoryLog(c *gin.Context) {
	id := c.Param("id")
	var log models.InventoryLog
	if err := database.DB.First(&log, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory log not found"})
		return
	}

	if err := database.DB.Delete(&log).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete inventory log"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory log deleted successfully"})
}
