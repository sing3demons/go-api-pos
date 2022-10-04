package main

import (
	"log"
	"os"

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
	db.ConnectDB()
	os.MkdirAll("uploads/products", 0755)
	
	r := gin.Default()
	r.Static("/uploads", "./uploads")
	serveRoutes(r)

	r.Run(":8080")
}
