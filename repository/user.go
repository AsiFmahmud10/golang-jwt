package repository

import (
	"fmt"
	"github.com/AsiFmahmud10/golang-jwt/config"
	"github.com/AsiFmahmud10/golang-jwt/model"
)

func StoreUser(newUser *model.User)*model.User{
	db := config.GetDb()
	db.Create(&model.User{Username:newUser.Username,Password:newUser.Password})
	var  new model.User
	db.First(&new)
	fmt.Print(new) 
	return &new 
}