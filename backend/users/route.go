package users

import (
	"pro-tasker/users/controller"

	"github.com/gin-gonic/gin"
)

func UsersRoutes(r *gin.RouterGroup) {
	// Define routes specific to package 1 within the specified router group
	r.GET("/list/:userID/", controller.ListUsers)
}
