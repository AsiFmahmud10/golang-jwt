package routes
import (
	"github.com/gorilla/mux"
	"github.com/AsiFmahmud10/golang-jwt/controllers"
)


func RegisterBookRoutes( r *mux.Router ){

	r.HandleFunc("/", controllers.SignUp).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
    r.HandleFunc("/welcome", controllers.Welcome).Methods("GET")
}