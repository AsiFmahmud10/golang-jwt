package controllers

import (
	"encoding/json"
	"github.com/AsiFmahmud10/golang-jwt/model"
	"github.com/AsiFmahmud10/golang-jwt/service"
	"fmt"
	"net/http"
	"time"
)

var key = []byte("my_secret_key")


func Login(w http.ResponseWriter,r *http.Request){
	var cred model.User
	json.NewDecoder(r.Body).Decode(&cred)
	fmt.Println("Request credi post : " + cred.Username)
	user, statusCode := service.Auth(&cred)
	if(statusCode != http.StatusOK){
			w.WriteHeader(http.StatusUnauthorized)
			return  
	}
		
	
	tokenString, err := service.Authorize(user,key)
	 if err != nil {
 		fmt.Print(err)
	 }

	http.SetCookie(w,&http.Cookie{
		Name: "token",
		Value: tokenString,
		Expires: time.Now().Add(1* time.Hour),
	})
		
}