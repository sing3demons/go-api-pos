package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/sing3demons/pos-app/db"
	"github.com/sing3demons/pos-app/dto"
	"github.com/sing3demons/pos-app/model"
)

type Order struct{}

func (order Order) FindAll(ctx *gin.Context) {
	var orders []model.Order
	db.Conn.Preload("Products").Find(&orders)

	result := []dto.OrderResponse{}

	for _, order := range orders {
		orderResult := dto.OrderResponse{
			ID:    order.ID,
			Name:  order.Name,
			Tel:   order.Tel,
			Email: order.Email,
		}
		products := []dto.OrderProductResponse{}
		for _, product := range order.Products {
			products = append(products, dto.OrderProductResponse{
				ID:       product.ID,
				SKU:      product.SKU,
				Name:     product.Name,
				Price:    product.Price,
				Quantity: product.Quantity,
				Image:    product.Image,
			})
		}
		orderResult.Products = products
		result = append(result, orderResult)
	}

	ctx.JSON(http.StatusOK, result)
}

func (o Order) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	var order model.Order

	query := db.Conn.Preload("Products").First(&order, id)
	if err := query.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	result := dto.OrderResponse{
		ID:    order.ID,
		Name:  order.Name,
		Tel:   order.Tel,
		Email: order.Email,
	}

	var products []dto.OrderProductResponse
	for _, product := range order.Products {
		products = append(products, dto.OrderProductResponse{
			ID:       product.ID,
			SKU:      product.SKU,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
			Image:    product.Image,
		})
	}

	result.Products = products

	ctx.JSON(http.StatusOK, result)
}

func (o Order) Create(ctx *gin.Context) {
	var body dto.OrderRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var order model.Order
	var orderItems []model.OrderItem

	for _, product := range body.Products {
		orderItems = append(orderItems, model.OrderItem{
			SKU:      product.SKU,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
			Image:    product.Image,
		})
	}

	order.Name = body.Name
	order.Tel = body.Tel
	order.Email = body.Email
	order.Products = orderItems

	db.Conn.Create(&order)

	result := dto.OrderResponse{
		ID:    order.ID,
		Name:  order.Name,
		Tel:   order.Tel,
		Email: order.Email,
	}

	var products []dto.OrderProductResponse
	for _, product := range order.Products {
		products = append(products, dto.OrderProductResponse{
			ID:       product.ID,
			SKU:      product.SKU,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
			Image:    product.Image,
		})
	}
	result.Products = products

	ctx.JSON(http.StatusCreated, result)
}
