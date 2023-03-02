package service
import(
	"github.com/AsiFmahmud10/golang-jwt/model"
	"github.com/AsiFmahmud10/golang-jwt/repository"

)

func RegisterUser(newUser*model.User)(*model.User,error){
	new :=repository.StoreUser(newUser)
	return new, nil
}