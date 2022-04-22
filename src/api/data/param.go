package data

import "github.com/gin-gonic/gin"

func BindParam(c *gin.Context) {
	var customer Customer
	if err := c.ShouldBindUri(&customer); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"name": customer.Name, "address": customer.Address, "birthdate": customer.Birthdate})
}

type Customer struct {
	Name      string `uri:"name" binding:"required"`
	Address   string `uri:"address" binding:"required"`
	Birthdate string `uri:"birthdate" binding:"required"`
}
