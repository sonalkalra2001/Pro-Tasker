package main

import (
	"fmt"
	"pro-tasker/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// var (
	// 	dbConn *gorm.DB
	// 	err    error
	// )
	// fmt.Println("Welcome To Pro - Tasker ::) ")
	// if dbConn, err = lib.DbConnection(lib.ConnDet{
	// 	Tx: true,
	// }); err != nil {
	// 	return
	// }
	// // begin a transaction

	// test := db.Test{
	// 	Name: "sonalika",
	// }
	// if result := dbConn.Create(&test); result.Error != nil {
	// 	log.Fatal("err in insertion")
	// 	dbConn.Rollback()
	// } else {
	// 	fmt.Println("successfully inserted")
	// 	dbConn.Commit()
	// }

	// Initialize the Gin router
	router := gin.Default()

	// Set up all routes for the application
	routes.SetupRoutes(router)

	// Start the server on port 8080
	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	router.Run(fmt.Sprintf(":%d", port))

}
