package routers

import (
	"gin-starter/src/api/ascii"
	"gin-starter/src/api/asyncsync"
	"gin-starter/src/api/bookable"
	"gin-starter/src/api/data"
	"gin-starter/src/api/ping"
	"gin-starter/src/api/queryMap"
	"gin-starter/src/api/validators"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Router(r *gin.Engine) *gin.Engine {
	// Register Validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookableValidator", validators.BookableValidator)
	}
	// Test API
	r.GET("/", func(c *gin.Context) { time.Sleep(5 * time.Second); c.String(http.StatusOK, "Welcome Gin Server") })
	r.GET("/ping", ping.Pong)
	r.GET("/asciiJSON", ascii.AsciiJson)
	r.GET("/bindformb", data.GetDataB)
	r.GET("/bindformc", data.GetDataC)
	r.GET("/bindformd", data.GetDataD)
	r.GET("/bindQuery", data.BindQuery)
	r.GET("/bindParams/:name/:address/:birthdate", data.BindParam)
	r.GET("/bookable", bookable.BookValidate)
	r.GET("/long_async", asyncsync.Async)
	r.GET("/long_sync", asyncsync.Sync)
	// API GROUP
	v1 := r.Group("/v1")
	{
		v1.POST("/login", func(ctx *gin.Context) { ctx.String(http.StatusOK, "login success") })
		v1.POST("/submit", func(ctx *gin.Context) { ctx.String(http.StatusOK, "submit success") })
		v1.POST("/read", func(ctx *gin.Context) { ctx.String(http.StatusOK, "read success") })
	}
	// HTML rendering
	r.GET("/index", func(c *gin.Context) { c.HTML(http.StatusOK, "welcome/index.tmpl", gin.H{"title": "main"}) })
	// Map as QueryString or Postform parameters
	r.POST("/queryMap", queryMap.MapToString)
	return r
}
