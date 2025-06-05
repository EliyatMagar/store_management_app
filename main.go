package main

import (
	"log"
	"os"
	"strings"
	"time"

	"store-app/config"
	"store-app/database"
	"store-app/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.Connect()

	router := gin.Default()

	// Load CORS origins from .env
	allowedOrigins := strings.Split(os.Getenv("CORS_ORIGIN"), ",")

	// Apply CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Register routes
	routes.RegisterAuthRoutes(router)
	routes.RegisterUserRoutes(router)
	routes.RegisterCategoryRoutes(router)
	routes.RegisterProductRoutes(router)
	routes.RegisterInventoryLogRoutes(router)
	routes.RegisterCustomerRoutes(router)
	routes.RegisterOrderRoutes(router)
	routes.RegisterOrderItemRoutes(router)
	routes.RegisterReviewRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(router.Run(":" + port))
}
