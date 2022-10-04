package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sing3demons/pos-app/dto"
)

type Category struct{}

func (c Category) FindAll(ctx *gin.Context) {}
func (c Category) FindOne(ctx *gin.Context) {}
func (c Category) Create(ctx *gin.Context) {

	var form dto.CategoryRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"Name": form.Name,
	})
}
func (c Category) Update(ctx *gin.Context) {}
func (c Category) Delete(ctx *gin.Context) {}
