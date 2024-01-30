package service

import (
	"pro-tasker/lib"
	"pro-tasker/users/mapper"
	"pro-tasker/users/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (reqModel *ListUsers) ProcessReq(c *gin.Context) (data interface{}, err error) {
	var (
		dbConn    *gorm.DB
		usersList []mapper.ListUsers
	)
	if dbConn, err = lib.DbConnection(lib.ConnDet{
		Writable: false,
	}); err != nil {
		return
	}
	if usersList, err = model.ListUsers(dbConn, reqModel.UserID); err != nil {
		return
	}
	data = usersList
	return
}
