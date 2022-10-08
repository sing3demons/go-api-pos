package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sing3demons/pos-app/db"
)

func init() {
	if os.Getenv("GIN_MODE") != gin.ReleaseMode {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db.ConnectDB()
	db.Migration()
	os.MkdirAll("uploads/products", 0755)

	r := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true

	r.Use(cors.New(corsConfig))
	r.Static("/uploads", "./uploads")
	serveRoutes(r)

	listenAndServe(r, port)
}

func listenAndServe(r *gin.Engine, port string) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	s := &http.Server{
		Addr:           ":" + port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	stop()
	fmt.Println("shutting down gracefully, press Ctrl+C again to force")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(timeoutCtx); err != nil {
		fmt.Println(err)
	}
}
