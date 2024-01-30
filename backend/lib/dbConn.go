package lib

import (
	"fmt"
	"log"
	"pro-tasker/db"
	"pro-tasker/lib/constant"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func DbConnection(connDet ConnDet) (dbConn *gorm.DB, err error) {
	var (
		dsn string
	)
	// Create a connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", constant.Host,
		constant.Port, constant.Username, constant.Password, constant.DbName)
	connStrSlave := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", constant.Host,
		constant.Port, constant.SlaveUsername, constant.SlavePassword, constant.DbName)

	if connDet.Tx || connDet.Writable {
		dsn = connStr
	} else { // readable conn
		dsn = connStrSlave
	}
	// Open a gorm connection
	if dbConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
		PrepareStmt: true, // caches so that queries can be speed up
	}); err != nil {
		log.Fatal("Error while opening a gorm connection", err)
		return
	}
	// test the connection
	if err = dbConn.Exec("SELECT 1").Error; err != nil {
		log.Fatal("Database connection failed:", err)
	}
	// Enable Gorm's logging
	dbConn.Logger = logger.Default.LogMode(logger.Info)

	// create a test table to check if db is working fine
	if err = dbConn.AutoMigrate(&db.Test{}); err != nil {
		return
	}
	// initialise a trnsaction
	if connDet.Tx {
		dbConn = dbConn.Begin()
	}
	return
}
