package service
import(
	"fmt"
	"net/http"
	"time"
	"github.com/AsiFmahmud10/golang-jwt/model"
	"github.com/AsiFmahmud10/golang-jwt/types"
	"github.com/AsiFmahmud10/golang-jwt/config"
	"github.com/golang-jwt/jwt/v4"
)


func Auth( cred *model.User )(*model.User, int){
	db := config.GetDb()
	var dbUser model.User
	db.First(&dbUser, "username = ?", cred.Username)
	if dbUser.Password != cred.Password{
		
		return nil,http.StatusUnauthorized
	}

	fmt.Println("saved user - " + dbUser.Username)

	return &dbUser , http.StatusOK
}


func Authorize( validUser *model.User, key []byte )(string, error){

	claims := types.Claims {
		validUser.Username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 *time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	tokenString, err := token.SignedString(key)
	
	if err != nil {
		fmt.Print(err)
	}

	return  tokenString,err


}