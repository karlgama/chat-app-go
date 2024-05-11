package user_use_cases

import (
	"github.com/karlgama/chat-app-go.git/domain/entities"
	security "github.com/karlgama/chat-app-go.git/infra/security/services"
	"github.com/sirupsen/logrus"
)

type CreateUserInput struct {
	Name     string `json:"name" binding:"required,gte=2,"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=8"`
}

type CreateUserUseCase struct {
}

func CreateUser(input *CreateUserInput) (*entities.User, error) {
	hashedPassword, err := security.HashPassword(input.Password)
	logrus.Info("creating user")

	if err != nil {
		return nil, err
	}

	return entities.NewUser(
		input.Name,
		input.Email,
		hashedPassword,
	), nil
}
