package controllers
import(
	"net/http"
	"encoding/json"
	"github.com/AsiFmahmud10/golang-jwt/service"
	"github.com/AsiFmahmud10/golang-jwt/model"
)

func SignUp(w http.ResponseWriter,r *http.Request){
	var cred  model.User
	json.NewDecoder(r.Body).Decode(&cred)
	_, err := service.RegisterUser(&cred)
	
	if  err!= nil{
		w.WriteHeader(401)
		return
	}
		w.WriteHeader(200)
		return 
}
