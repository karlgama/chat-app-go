package user_use_cases

import (
	"github.com/karlgama/chat-app-go.git/domain/entities"
	security "github.com/karlgama/chat-app-go.git/infra/security/services"
)

func FindUserByEmail(email string) *entities.User {
	hashedPassword, _ := security.HashPassword("12345678")

	return &entities.User{
		ID:         nil,
		ExternalID: nil,
		Name:       "name",
		Email:      email,
		Password:   hashedPassword,
		CreatedAt:  nil,
		UpdatedAt:  nil,
	}
}
