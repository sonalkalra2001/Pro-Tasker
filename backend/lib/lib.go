package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendAPIResp(data interface{}, err error, c *gin.Context) {

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": data,
		})
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}
