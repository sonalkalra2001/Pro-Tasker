package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendAPIResp(data ApiResp, err error, c *gin.Context) {

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":    "success",
			"data":      data.Data,
			"debug_msg": data.DebugMsg,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":     err.Error(),
			"status":    "error",
			"debug_msg": data.DebugMsg,
		})
	}
}
