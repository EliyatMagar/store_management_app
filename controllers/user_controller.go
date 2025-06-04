package controllers

import (
	"net/http"
	"store-app/database"
	"store-app/models"

	"github.com/gin-gonic/gin"
)

// PublicUser is a safe version of the User struct for API responses.
type PublicUser struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

// convertToPublicUser maps a models.User to PublicUser
func convertToPublicUser(user models.User) PublicUser {
	return PublicUser{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
}

// Get all users
func GetAllUsers(c *gin.Context) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	publicUsers := make([]PublicUser, 0, len(users))
	for _, user := range users {
		publicUsers = append(publicUsers, convertToPublicUser(user))
	}

	c.JSON(http.StatusOK, gin.H{"users": publicUsers})
}

// Get user by ID
func GetUserByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": convertToPublicUser(user)})
}

// Update user by ID
func UpdateUserByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Role  string `json:"role"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Role = input.Role

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": convertToPublicUser(user)})
}

// Delete user by ID
func DeleteUserByID(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
