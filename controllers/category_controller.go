package controllers

import (
	"net/http"
	"sort"
	"store-app/database"
	"store-app/models"

	"github.com/gin-gonic/gin"
)

// Create a new category
func CreateCategory(c *gin.Context) {
	var input struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := models.Category{
		Name:        input.Name,
		Description: input.Description,
	}

	if err := database.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully", "category": category})
}

// Get all categories (with optional alphabetical sort)
func GetAllCategories(c *gin.Context) {
	var categories []models.Category

	if err := database.DB.Preload("Products").Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve categories"})
		return
	}

	// Sort alphabetically by category name
	sort.SliceStable(categories, func(i, j int) bool {
		return categories[i].Name < categories[j].Name
	})

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

// Get a single category by ID
func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	if err := database.DB.Preload("Products").First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"category": category})
}

// Update a category by ID
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	var input struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.Name = input.Name
	category.Description = input.Description

	if err := database.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully", "category": category})
}

// Delete a category by ID
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	if err := database.DB.Delete(&models.Category{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
