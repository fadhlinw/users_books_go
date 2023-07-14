package config

import (
	"tugas/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const DSN = "root:12345@tcp(localhost:3306)/tugas?charset=utf8&parseTime=True&loc=Local"

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitialMigration()
}

func InitialMigration() {
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.User{})
}
