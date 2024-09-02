package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func Connect(models []interface{},path string){
	database, err := gorm.Open(sqlite.Open(path), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database!")
	}

	err = database.AutoMigrate(models...)
	if err != nil {
		panic("Error migrating database!")
	}

	Conn = database
}