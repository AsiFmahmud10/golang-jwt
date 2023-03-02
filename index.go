package main
import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/AsiFmahmud10/golang-jwt/config"
	"github.com/AsiFmahmud10/golang-jwt/routes"

)
var key = []byte("my_secret_key")

var users = map[string]string{
	"u1" : "p1",
} 


func main(){
	r := mux.NewRouter()
	config.Connect()
	routes.RegisterBookRoutes(r)
	http.ListenAndServe(":80",r)
}

