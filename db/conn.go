package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Conn *gorm.DB

func ConnectDB() {
	dsn := "kpsing:P@ssw0rd@tcp(127.0.0.1:3306)/go-pos?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	Conn = db
}