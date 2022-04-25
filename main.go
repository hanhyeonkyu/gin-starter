package main

import (
	"context"
	"gin-starter/src/api/logger"
	"gin-starter/src/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	PORT := "8080"
	gin.ForceConsoleColor()

	// Log To File (if you use it your log don't show in your console)
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)
	r1 := gin.Default()
	// serve static files
	r1.Static("/assets", "./assets")
	r1.StaticFS("/more_static", http.Dir("my_file_system"))
	r1.StaticFile("/favicon.ico", "./resources/favicon.ico")
	// Custum Log (upper line remove then replace down line)
	r2 := gin.New()
	r2 = logger.Logger(r2)
	// second http handler
	r2.Use(gin.Logger())
	r2.Use(gin.Recovery())
	r2.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":    http.StatusOK,
				"message": "Welcome server 01",
			},
		)
	})

	// Load HTML Template with Glob
	r1.LoadHTMLGlob("templates/**/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	// Get Routes
	r1 = routers.Router(r1)

	s1 := &http.Server{
		Addr:           ":" + PORT,
		Handler:        r1,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s2 := &http.Server{
		Addr:         ":8081",
		Handler:      r2,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := s2.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// graceful shutdown
	go func() {
		if err := s1.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdonw Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s1.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdon: ", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds")
	}
	log.Println("Server exiting")
}
