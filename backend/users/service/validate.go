package service

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (reqModel *ListUsers) ValidateReq(req *gin.Context) (err error) {

	if reqModel.UserID, err = strconv.Atoi(req.Param("userID")); err != nil {
		return
	}
	if reqModel.UserID == 0 {
		fmt.Println("Please provide a valid user ID ")
	}
	return
}
