package queryMap

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MapToString(c *gin.Context) {
	ids := c.QueryMap("ids")
	names := c.PostFormMap("names")
	b := new(bytes.Buffer)
	for k, v := range ids {
		fmt.Fprintf(b, "%s=\"%s\"\n", k, v)
	}
	for k, v := range names {
		fmt.Fprintf(b, "%s=\"%s\"\n", k, v)
	}
	c.String(http.StatusOK, b.String())
}
