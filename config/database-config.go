package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var db *gorm.DB

func ConnectDB() (db *gorm.DB, err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/test.db?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func InitConnection() {
	dbValue, err := ConnectDB()

	db = dbValue

	if err != nil {
		fmt.Printf("there was something wrong to connect DB: %s", err.Error())
	}
}

func GetDB() *gorm.DB {
	return db
}
