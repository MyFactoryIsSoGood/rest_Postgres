package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 user=postgres password=123qwe123 dbname=movies sslmode=disable")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Movie{})

	DB = db
}
