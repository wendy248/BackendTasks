package config

import (
	"Github/BackendTasks/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@(localhost)/dbbackend?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("connection to database error")
	}

	db.AutoMigrate(&models.Animal{})
	return db
}
