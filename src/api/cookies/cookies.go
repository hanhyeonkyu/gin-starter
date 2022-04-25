package cookies

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GinCookie(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")
	if err != nil {
		cookie = "Not Set"
		c.SetCookie("gin_cookie", "test", 10, "/", "localhost", false, true)
	}
	fmt.Printf("Cookie value: %s \n", cookie)
}
