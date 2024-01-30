package routes

import (
	"pro-tasker/users"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(c *gin.Engine) {
	proTasker := c.Group("/proTasker")
	users.UsersRoutes(proTasker.Group("/users"))

}
