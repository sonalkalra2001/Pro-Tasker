package main

import (
	"fmt"
	"log"
	"pro-tasker/db"
	"pro-tasker/lib"

	"gorm.io/gorm"
)

func main() {
	var (
		dbConn *gorm.DB
		err    error
	)
	fmt.Println("Welcome To Pro - Tasker ::) ")
	if dbConn, err = lib.DbConnection(lib.ConnDet{
		Tx: true,
	}); err != nil {
		return
	}
	// begin a transaction

	test := db.Test{
		Name: "sonalika",
	}
	if result := dbConn.Create(&test); result.Error != nil {
		log.Fatal("err in insertion")
		dbConn.Rollback()
	} else {
		fmt.Println("successfully inserted")
		dbConn.Commit()
	}

	fmt.Println(dbConn)

}
