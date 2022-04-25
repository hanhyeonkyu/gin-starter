package upload

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Multi(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["upload"]
	for _, file := range files {
		log.Println(file.Filename)
		c.SaveUploadedFile(file, "./"+file.Filename)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

func Single(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	c.SaveUploadedFile(file, "./"+file.Filename)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
