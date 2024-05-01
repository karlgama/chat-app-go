package use_cases

import (
	"github.com/karlgama/chat-app-go.git/domain/entities"
)

type CreateUserInput struct {
	Name     string `json:"name" validate:"required,min=3,"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func CreateUser(input *CreateUserInput) *entities.User {

	return entities.NewUser(
		input.Name,
		input.Email,
		input.Password,
	)
}
