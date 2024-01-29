package lib

import (
	"fmt"
	"log"
	"pro-tasker/db"
	"pro-tasker/lib/constant"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection(connDet ConnDet) (dbConn *gorm.DB, err error) {
	// Create a connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", constant.Host,
		constant.Port, constant.Username, constant.Password, constant.DbName)

	// Open a gorm connection
	if dbConn, err = gorm.Open(postgres.Open(connStr), &gorm.Config{}); err != nil {
		log.Fatal("Error while opening a gorm connection", err)
		return
	}
	// test the connection
	if err = dbConn.Exec("SELECT 1").Error; err != nil {
		log.Fatal("Database connection failed:", err)
	}
	// create a test table to check if db is working fine
	dbConn.AutoMigrate(&db.Test{})
	if connDet.Tx {
		dbConn = dbConn.Begin()
	}
	return
}
