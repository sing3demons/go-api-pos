package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	os.MkdirAll("uploads/products", 0755)
	r := gin.Default()
	r.Static("/uploads", "./uploads")
	serveRoutes(r)

	r.Run(":8080")
}
