package model

import (
	"pro-tasker/db"
	"pro-tasker/users/mapper"

	"gorm.io/gorm"
)

func ListUsers(dbConn *gorm.DB, userID int) (usersList []mapper.ListUsers, err error) {

	if result := dbConn.Model(&db.Test{}).Select("name , test_id AS user_id ").Where("test_id >0 ").
		Scan(&usersList); result.Error != nil {
		return
	}
	return
}
