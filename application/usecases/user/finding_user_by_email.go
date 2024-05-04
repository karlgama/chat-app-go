package user_use_cases

import (
	"github.com/karlgama/chat-app-go.git/domain/entities"
	"github.com/karlgama/chat-app-go.git/infra/security"
)

func FindUserByEmail(email string) *entities.User {
	hashedPassword, _ := security.HashPassword("12345678")

	return &entities.User{
		nil,
		"",
		email,
		hashedPassword,
		nil,
		nil,
	}
}
