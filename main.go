package main

import (
	"log"
	"os"

	"store-app/config"
	"store-app/database"
	"store-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.Connect()

	router := gin.Default()
	// Register routes
	routes.RegisterAuthRoutes(router)
	routes.RegisterUserRoutes(router)
	routes.RegisterCategoryRoutes(router)
	routes.RegisterProductRoutes(router)
	routes.RegisterInventoryLogRoutes(router)
	routes.RegisterCustomerRoutes(router)
	routes.RegisterOrderRoutes(router)
	routes.RegisterOrderItemRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(router.Run(":" + port))
}
