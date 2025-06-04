package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"store-app/database"
	"store-app/models"
)

// GetCustomers retrieves all customers from the database

func GetCustomers(c *gin.Context) {
	var customers []models.Customer
	if err := database.DB.Find(&customers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve customers"})
		return
	}

	c.JSON(http.StatusOK, customers)
}

func GetCustomerByID(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer
	if err := database.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := database.DB.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	c.JSON(http.StatusCreated, customer)

}

func CreateMultipleCustomers(c *gin.Context) {
	var customers []models.Customer

	if err := c.ShouldBindJSON(&customers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := database.DB.Create(&customers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customers"})
		return
	}

	c.JSON(http.StatusOK, customers)
}

func UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer
	if err := database.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	var updatedCustomer models.Customer
	if err := c.ShouldBindJSON(&updatedCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	database.DB.Model(&customer).Updates(updatedCustomer)
	c.JSON(http.StatusOK, customer)

}

func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")

	if err := database.DB.Delete(&models.Customer{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "Customer deleted successfully"})
}
