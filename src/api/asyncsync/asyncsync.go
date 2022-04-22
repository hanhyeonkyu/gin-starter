package asyncsync

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Async(c *gin.Context) {
	cCp := c.Copy()
	go func() {
		time.Sleep(1 * time.Second)
		log.Println("Done! in path " + cCp.Request.URL.Path)
		http.Get("http://localhost:8080/long_sync")
	}()

}

func Sync(c *gin.Context) {
	time.Sleep(1 * time.Second)
	http.Get("http://localhost:8080/long_async")
	log.Println("Done! in path " + c.Request.URL.Path)
}
