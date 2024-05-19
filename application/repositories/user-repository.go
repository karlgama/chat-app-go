package repositories

import "github.com/karlgama/chat-app-go.git/domain/entities"

type UserRepository interface {
	Save(user *entities.User) (*entities.User, error)
}
