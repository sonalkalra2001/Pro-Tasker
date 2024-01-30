package controller

import (
	"pro-tasker/lib"
	"pro-tasker/users/service"

	"github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {

	var (
		reqModel service.ListUsers
		err      error
		data     interface{}
	)

	if err = (&reqModel).ValidateReq(c); err == nil {
		if data, err = (&reqModel).ProcessReq(c); err == nil {
			// send response
			lib.SendAPIResp(data, err, c)
		}
	}
}
