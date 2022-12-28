package httphandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetJson(obj any, c *gin.Context) bool {

	if err := c.ShouldBindJSON(&obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}

	return true
}
