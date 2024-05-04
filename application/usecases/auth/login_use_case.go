package auth_use_cases

import (
	"errors"

	user_use_cases "github.com/karlgama/chat-app-go.git/application/usecases/user"
	security "github.com/karlgama/chat-app-go.git/infra/security/services"
	"github.com/sirupsen/logrus"
)

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

	token, err := security.GenerateToken(foundUser)

	if err != nil {
		return nil, errors.New("could not generate token")
	}
	return &LoginOutput{Token: &token}, nil
}
