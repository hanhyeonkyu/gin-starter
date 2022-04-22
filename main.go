package main

import (
	"context"
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

	r := gin.Default()
	// Custum Log (upper line remove then replace down line)
	// r := gin.New()
	// r = logger.Logger(r)

	// Load HTML Template with Glob
	r.LoadHTMLGlob("templates/**/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	// Get Routes
	r = routers.Router(r)

	s := &http.Server{
		Addr:           ":" + PORT,
		Handler:        r,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// graceful shutdown
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdonw Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdon: ", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds")
	}
	log.Println("Server exiting")
}
