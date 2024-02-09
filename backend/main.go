package main

import (
	"fmt"
	"pro-tasker/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize the Gin router
	router := gin.Default()

	// Set up all routes for the application
	routes.SetupRoutes(router)

	// Start the server on port 8080
	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	router.Run(fmt.Sprintf(":%d", port))

}
