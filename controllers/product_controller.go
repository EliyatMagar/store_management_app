package controllers

import (
	"fmt"
	"net/http"
	"store-app/database"
	"store-app/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

//GenerateSKU creates a basic SKU using product name and timestamp

func GenerateSKU(name string) string {
	timestamp := time.Now().Unix()
	processed := strings.ToUpper(strings.ReplaceAll(name, " ", "_"))
	if len(processed) > 10 {
		processed = processed[:10]
	}

	return fmt.Sprintf("%s_%d", processed, timestamp)
}

// CreateProduct creates a new product with category_id or category_name support
func CreateProduct(c *gin.Context) {
	var input struct {
		Name         string  `json:"name" binding:"required"`
		Description  string  `json:"description"`
		Price        float64 `json:"price" binding:"required"`
		Stock        int     `json:"stock" binding:"required"`
		CategoryID   uint    `json:"category_id"`   // Optional
		CategoryName string  `json:"category_name"` // Optional
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var category models.Category

	// Check if either category_id or category_name is provided
	if input.CategoryID != 0 {
		// Lookup by ID
		if err := database.DB.First(&category, input.CategoryID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found by ID"})
			return
		}
	} else if input.CategoryName != "" {
		// Lookup by name
		if err := database.DB.Where("name = ?", input.CategoryName).First(&category).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found by name"})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Either category_id or category_name is required"})
		return
	}

	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
		CategoryID:  category.ID,
		SKU:         GenerateSKU(input.Name),
	}

	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
		"product": product,
	})
}

// CreateMultipleProducts allows batch creation of products with category_name or category_id
func CreateMultipleProducts(c *gin.Context) {
	var inputs []struct {
		Name         string  `json:"name" binding:"required"`
		Description  string  `json:"description"`
		Price        float64 `json:"price" binding:"required"`
		Stock        int     `json:"stock" binding:"required"`
		CategoryID   uint    `json:"category_id"`
		CategoryName string  `json:"category_name"`
	}

	if err := c.ShouldBindJSON(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var createdProducts []models.Product

	for _, input := range inputs {
		var category models.Category

		if input.CategoryID != 0 {
			if err := database.DB.First(&category, input.CategoryID).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Category not found by ID", "product": input.Name})
				return
			}
		} else if input.CategoryName != "" {
			if err := database.DB.Where("name = ?", input.CategoryName).First(&category).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Category not found by name", "product": input.Name})
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Either category_id or category_name is required", "product": input.Name})
			return
		}

		product := models.Product{
			Name:        input.Name,
			Description: input.Description,
			Price:       input.Price,
			Stock:       input.Stock,
			CategoryID:  category.ID,
			SKU:         GenerateSKU(input.Name),
		}

		if err := database.DB.Create(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product", "product": input.Name})
			return
		}

		createdProducts = append(createdProducts, product)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Products created successfully",
		"products": createdProducts,
	})
}

//GetAllProducts fetches all products

func GetAllProducts(c *gin.Context) {
	var products []models.Product
	if err := database.DB.Preload("Category").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

// GetProductByID fetches a product by ID
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := database.DB.Preload("Category").First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}

//UpdateProduct updates product info

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var input struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Stock       int     `json:"stock"`
		CategoryID  int     `json:"category_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Name != "" {
		product.Name = input.Name
		product.SKU = GenerateSKU(input.Name) // Regenerate SKU if name changes
	}

	if input.Description != "" {
		product.Description = input.Description
	}
	if input.Price != 0 {
		product.Price = input.Price
	}
	if input.Stock != 0 {
		product.Stock = input.Stock
	}
	if input.CategoryID != 0 {
		var category models.Category
		if err := database.DB.First(&category, input.CategoryID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}
		product.CategoryID = uint(input.CategoryID)
	}

	if err := database.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully", "product": product})
}

//DeleteProduct removes a product

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	if err := database.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

//SearchProducts searches products by name or description

func SearchProducts(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	var products []models.Product
	if err := database.DB.Preload("Category").Where("name LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

// GetProductsByCategory fetches products by category ID
func GetProductsByCategory(c *gin.Context) {
	categoryID := c.Param("category_id")
	var products []models.Product

	if err := database.DB.Preload("Category").Where("category_id = ?", categoryID).Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products by category"})
		return
	}

	if len(products) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No products found for this category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

// GetProductStock fetches stock info for a product
func GetProductStock(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product_id": product.ID, "stock": product.Stock})
}
