package entities

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	jwt.StandardClaims
}
