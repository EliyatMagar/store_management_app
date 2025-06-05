package controllers

import (
	"log"
	"net/http"
	"store-app/database"
	"store-app/models"

	"github.com/gin-gonic/gin"
)

func CreateReview(c *gin.Context) {
	var review models.Review

	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&review).Error; err != nil {
		log.Println("DB Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Make sure ProductID and CustomerID exist"})
		return
	}

	c.JSON(http.StatusCreated, review)
}

func GetAllReviews(c *gin.Context) {
	var reviews []models.Review
	if err := database.DB.Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reviews"})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

func GetReviewByID(c *gin.Context) {
	id := c.Param("id")
	var review models.Review
	if err := database.DB.First(&review, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		return
	}
	c.JSON(http.StatusOK, review)
}

func UpdateReview(c *gin.Context) {
	id := c.Param("id")
	var review models.Review

	if err := database.DB.First(&review, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		return
	}

	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update review"})
		return
	}
	c.JSON(http.StatusOK, review)
}

func DeleteReview(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Review{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete review"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}
