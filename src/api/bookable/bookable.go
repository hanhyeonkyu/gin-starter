package bookable

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func BookValidate(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Book is Valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
}

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookableValidator" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn,bookableValidator" time_format:"2006-01-02"`
}
