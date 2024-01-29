package main

import (
	"fmt"
	"pro-tasker/lib"

	"gorm.io/gorm"
)

func main() {
	var (
		dbConn *gorm.DB
		err    error
	)
	fmt.Println("Welcome To Pro - Tasker ::) ")
	if dbConn, err = lib.DbConnection(); err != nil {
		return
	}
	fmt.Println(dbConn)

}
