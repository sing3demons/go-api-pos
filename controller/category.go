package controller

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sing3demons/pos-app/db"
	"github.com/sing3demons/pos-app/dto"
	"github.com/sing3demons/pos-app/model"
	"gorm.io/gorm"
)

type Category struct{}

func (c Category) FindAll(ctx *gin.Context) {
	var categories []model.Category
	db.Conn.Find(&categories)

	result := []dto.CategoryResponse{}
	for _, category := range categories {
		result = append(result, dto.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (c Category) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	var category model.Category
	if err := db.Conn.First(&category, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	})
}

func (c Category) Create(ctx *gin.Context) {

	var form dto.CategoryRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := model.Category{Name: form.Name}

	if err := db.Conn.Save(&category).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	})
}

func (c Category) Update(ctx *gin.Context) {
	var form dto.CategoryRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := findCategoryByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	category.Name = form.Name
	if err := db.Conn.Save(&category).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	})
}

func (c Category) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := db.Conn.Delete(&model.Category{}, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"deletedAt": time.Now(),
	})
}

func findCategoryByID(ctx *gin.Context) (*model.Category, error) {
	id := ctx.Param("id")
	var category model.Category
	if err := db.Conn.First(&category, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &category, nil
}
