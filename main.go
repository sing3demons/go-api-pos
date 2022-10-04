package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sing3demons/pos-app/db"
)

func init() {
	db.ConnectDB()
}

func main() {

	os.MkdirAll("uploads/products", 0755)
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Static("/uploads", "./uploads")
	serveRoutes(r)

	r.Run(":8080")
}
