package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sing3demons/pos-app/controller"
)

func serveRoutes(r *gin.Engine) {
	productGroup := r.Group("/products")
	productController := controller.Product{}
	{
		productGroup.GET("", productController.FindAll)
		productGroup.GET("/:id", productController.FindOne)
		productGroup.POST("", productController.Create)
		productGroup.PATCH("/:id", productController.Update)
		productGroup.DELETE("/:id", productController.Delete)
	}
}
