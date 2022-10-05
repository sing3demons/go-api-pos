package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sing3demons/pos-app/dto"
)

type Product struct{}

func (p Product) FindAll(ctx *gin.Context) {
	search := ctx.Query("search")
	categoryId := ctx.Query("categoryId")

	ctx.JSON(http.StatusOK, gin.H{
		"FindAll":    "OK",
		"Search":     search,
		"CategoryID": categoryId,
	})
}
func (p Product) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"FindOne": "OK",
		"ID":      id,
	})
}
func (p Product) Create(ctx *gin.Context) {
	var form dto.ProductRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imageFile, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileName := strings.Split(imageFile.Filename, ".")[1]
	imagePath := "./uploads/products/" + uuid.New().String() + "." + fileName
	if err := ctx.SaveUploadedFile(imageFile, imagePath); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"Name": form.Name,
	})
}
func (p Product) Update(ctx *gin.Context) {}
func (p Product) Delete(ctx *gin.Context) {}
