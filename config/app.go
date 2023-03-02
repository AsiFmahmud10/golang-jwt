package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	_ "github.com/mattn/go-sqlite3"
	"github.com/AsiFmahmud10/golang-jwt/model"
)

var db *gorm.DB

func  GetDb() *gorm.DB {
	return db
}

func Connect(){
	d, err := gorm.Open(sqlite.Open("./repository/test.db"), &gorm.Config{})
	if err != nil{
		panic("Connection problem")
	}
	db = d
	db.AutoMigrate(&model.User{})
}