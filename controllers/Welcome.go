package controllers
import (
	"github.com/AsiFmahmud10/golang-jwt/types"
	"net/http"
	"github.com/golang-jwt/jwt/v4"
)


func Welcome(w http.ResponseWriter,r *http.Request){
	c , err:= r.Cookie("token")
	if err !=   nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	tknStr := c.Value
	claims := &types.Claims{}	
	tkn, err:= jwt.ParseWithClaims(tknStr,claims,func(t *jwt.Token) (interface{}, error) {
		return key,nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
	}

	if ! tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Write([]byte("welcome"))
}
