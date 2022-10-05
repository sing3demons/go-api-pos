package controller

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sing3demons/pos-app/db"
	"github.com/sing3demons/pos-app/dto"
	"github.com/sing3demons/pos-app/model"
	"gorm.io/gorm"
)

type Product struct{}

func (p Product) FindAll(ctx *gin.Context) {
	search := ctx.Query("search")
	categoryId := ctx.Query("categoryId")
	status := ctx.Query("status")

	// db.Conn.Where("category_id = ?",categoryId).Find(&products)

	var products []model.Product
	query := db.Conn.Preload("Category")
	if categoryId != "" {
		query = query.Where("category_id = ?", categoryId)
	}
	if search != "" {
		query = query.Where("sku = ? OR name LIKE ?", search, "%"+search+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Order("created_at desc").Find(&products).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	result := []dto.ProductResponse{}
	for _, product := range products {
		result = append(result, dto.ProductResponse{
			ID: product.ID, SKU: product.SKU,
			Name:   product.Name,
			Desc:   product.Desc,
			Price:  product.Price,
			Status: product.Status,
			Image:  product.Image,
			Category: dto.CategoryResponse{
				ID:   product.CategoryID,
				Name: product.Category.Name,
			},
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"products": result,
	})
}
func (p Product) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	var product model.Product
	query := db.Conn.Preload("Category").First(&product, id)
	if err := query.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, dto.ProductResponse{
		ID:     product.ID,
		SKU:    product.SKU,
		Name:   product.Name,
		Desc:   product.Desc,
		Price:  product.Price,
		Status: product.Status,
		Image:  product.Image,
		Category: dto.CategoryResponse{
			ID:   product.CategoryID,
			Name: product.Category.Name,
		},
	})
}
func (p Product) Create(ctx *gin.Context) {
	var body dto.ProductRequest
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imageFile, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	path := "uploads/products/"
	os.MkdirAll(path, 0755)

	fileName := strings.Split(imageFile.Filename, ".")[1]
	imagePath := path + uuid.New().String() + "." + fileName

	if err := ctx.SaveUploadedFile(imageFile, imagePath); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := model.Product{
		SKU:        body.SKU,
		Name:       body.Name,
		Desc:       body.Desc,
		Price:      body.Price,
		Status:     body.Status,
		CategoryID: body.CategoryID,
		Image:      imagePath,
	}

	if err := db.Conn.Save(&product).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateProductResponse{
		ID: product.ID, SKU: product.SKU,
		Name: product.Name, Desc: product.Desc,
		Price: product.Price, Status: product.Status,
		Image: product.Image, CategoryID: product.CategoryID,
	})
}
func (p Product) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var body dto.ProductRequest
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product model.Product
	if err := db.Conn.First(&product, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	imageFile, err := ctx.FormFile("image")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if imageFile != nil {
		fileName := strings.Split(imageFile.Filename, ".")[1]
		imagePath := "uploads/products/" + uuid.New().String() + "." + fileName
		if err := ctx.SaveUploadedFile(imageFile, imagePath); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		os.Remove(product.Image)
		product.Image = imagePath
	}

	product.SKU = body.SKU
	product.Name = body.Name
	product.Desc = body.Desc
	product.Price = body.Price
	product.Status = body.Status
	product.Status = body.Status
	product.CategoryID = body.CategoryID

	if err := db.Conn.Save(&product).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateProductResponse{
		ID:         product.ID,
		SKU:        product.SKU,
		Name:       product.Name,
		Desc:       product.Desc,
		Price:      product.Price,
		Status:     product.Status,
		Image:      product.Image,
		CategoryID: product.CategoryID,
	})
}

func (p Product) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	db.Conn.Delete(&model.Product{}, id)
	ctx.JSON(http.StatusOK, gin.H{
		"deletedAt": time.Now(),
	})
}
