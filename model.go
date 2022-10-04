package main

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;type:varchar(100);not null"`
}

type Order struct {
	gorm.Model
	Name     string
	Email    string
	Tel      string
	Products []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	gorm.Model
	SKU      string  `gorm:"not null"`
	Name     string  `gorm:"not null"`
	Image    string  `gorm:"not null"`
	Price    float64 `gorm:"not null"`
	Quantity uint    `gorm:"not null"`
	OrderID  uint
}

type Product struct {
	gorm.Model
	SKU        string `gorm:"uniqueIndex;type:varchar(100);not null"`
	Name       string `gorm:"type:varchar(100);not null"`
	Desc       string `gorm:"type:varchar(255)"`
	Price      float64
	Status     uint   `gorm:"not null"`
	Image      string `gorm:"type:varchar(255);not null"`
	CategoryID uint
	Category   Category
}
