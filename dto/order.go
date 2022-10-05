package dto

type OrderRequest struct {
	Name     string                `json:"name"`
	Tel      string                `json:"tel"`
	Email    string                `json:"email"`
	Products []OrderProductRequest `json:"products"`
}

type OrderProductRequest struct {
	Name     string  `json:"name"`
	SKU      string  `json:"sku"`
	Price    float64 `json:"price"`
	Status   uint    `json:"status"`
	Image    string  `json:"image"`
	Quantity uint    `json:"quantity"`
}

type OrderResponse struct {
	ID       uint                   `json:"id"`
	Name     string                 `json:"name"`
	Tel      string                 `json:"tel"`
	Email    string                 `json:"email"`
	Products []OrderProductResponse `json:"products"`
}

type OrderProductResponse struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	SKU      string  `json:"sku"`
	Price    float64 `json:"price"`
	Status   uint    `json:"status"`
	Image    string  `json:"image"`
	Quantity uint    `json:"quantity"`
}
