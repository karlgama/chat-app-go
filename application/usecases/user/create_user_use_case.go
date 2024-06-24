package user_use_cases

import (
	"github.com/karlgama/chat-app-go.git/application/repositories"
	"github.com/karlgama/chat-app-go.git/domain/entities"
	security "github.com/karlgama/chat-app-go.git/infra/security/services"
	"github.com/sirupsen/logrus"
)

type CreateUserInput struct {
	Name     string `json:"name" binding:"required,gte=2"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=8"`
}

type CreateUserUseCase struct {
	repository repositories.UserRepository
}

func (c *CreateUserUseCase) CreateUser(input *CreateUserInput) (*entities.User, error) {
	hashedPassword, err := security.HashPassword(input.Password)
	logrus.Info("creating user")

	if err != nil {
		return nil, err
	}

	user := entities.NewUser(
		input.Name,
		input.Email,
		hashedPassword,
	)

	savedUser, err := c.repository.Save(user)

	if err != nil {
		return nil, err
	}

	return savedUser, nil
}

func NewCreateUserUseCase(repository repositories.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{repository: repository}
}
