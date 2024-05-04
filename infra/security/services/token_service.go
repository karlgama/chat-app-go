package security_service

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/karlgama/chat-app-go.git/domain/entities"
)

var jwtKey = []byte("  ")

func GenerateToken(user *entities.User) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := entities.Claims{
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("could not parse claims")
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
