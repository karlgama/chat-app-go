package usecases

import (
	"github.com/karlgama/chat-app-go.git/application/DTOs/user/inputs"
	"github.com/karlgama/chat-app-go.git/domain/entities"
)

func CreateUser(input *inputs.CreateUserInput) {

	user := entities.User{
		input.Name,
		input.Email,
		input.Password,
	}
}
