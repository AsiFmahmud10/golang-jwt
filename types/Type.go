package types

import(
	"github.com/golang-jwt/jwt/v4"
)


type Cred struct{
	Username string	`json : 'username'`
	Password string `json : 'password'`
}

type Claims struct{
	Username string	`json : 'username'`;
	jwt.RegisteredClaims
}
