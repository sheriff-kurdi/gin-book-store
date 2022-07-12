package config

import (
	"github.com/rahmanfadhil/gin-bookstore/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnect() {
	dsn := "host=172.17.0.2 user=postgres password=StrongPassw0rd! dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Book{})
	DB = db
}
