package ascii

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AsciiJson(c *gin.Context) {
	data := map[string]interface{}{
		"lang": "GOLANG",
		"tag":  "hoho",
	}
	// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
	c.AsciiJSON(http.StatusOK, data)
}
