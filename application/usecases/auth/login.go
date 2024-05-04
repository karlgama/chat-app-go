package auth_use_cases

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	user_use_cases "github.com/karlgama/chat-app-go.git/application/usecases/user"
	"github.com/karlgama/chat-app-go.git/domain/entities"
	"github.com/karlgama/chat-app-go.git/infra/security"
	"github.com/sirupsen/logrus"
)

var jwtKey = []byte("  ")

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginOutput struct {
	Token *string `json:"token"`
}

func Login(input *LoginInput) (*LoginOutput, error) {

	foundUser := user_use_cases.FindUserByEmail(input.Email)
	logrus.Info("founduser", foundUser)

	if foundUser == nil {
		return nil, errors.New("email or password is incorrect")
	}

	errHash := security.CheckPasswordHash(input.Password, foundUser.Password)

	if !errHash {
		return nil, errors.New("email or password is incorrect")
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := entities.Claims{
		StandardClaims: jwt.StandardClaims{
			Subject:   foundUser.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return nil, errors.New("could not generate token")
	}
	return &LoginOutput{Token: &tokenString}, nil
}
