package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	c.Status(http.StatusOK)
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
