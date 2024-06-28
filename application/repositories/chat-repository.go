package repositories

import "github.com/karlgama/chat-app-go.git/domain/entities"

type ChatRepository interface {
	Save(*entities.Chat) (*entities.Chat, error)
}
