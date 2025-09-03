package postgreSQL

import (
	"github.com/google/uuid"
	"github.com/karlgama/chat-app-go.git/domain/entities"
	"github.com/karlgama/chat-app-go.git/infra/postgreSQL"
	"github.com/karlgama/chat-app-go.git/infra/postgreSQL/models"
)

type ChatPostgresRepository struct{}

func (c *ChatPostgresRepository) Save(chat *entities.Chat) (*entities.Chat, error) {
	// Gera um UUID para o chat se não existir
	if chat.ID == uuid.Nil {
		chat.ID = uuid.New()
	}

	chatModel := models.ChatModel{
		ID:        chat.ID,
		CreatedAt: chat.CreatedAt,
		UpdatedAt: chat.UpdatedAt,
	}

	// Salva o chat
	result := postgreSQL.DB.Create(&chatModel)
	if result.Error != nil {
		return nil, result.Error
	}

	// Se há usuários, cria as associações
	if chat.Users != nil && len(*chat.Users) > 0 {
		for _, user := range *chat.Users {
			if user.ID != nil {
				chatUserModel := models.ChatUserModel{
					ChatID: chat.ID,
					UserID: *user.ID,
				}
				postgreSQL.DB.Create(&chatUserModel)
			}
		}
	}

	// Atualiza os timestamps
	chat.CreatedAt = chatModel.CreatedAt
	chat.UpdatedAt = chatModel.UpdatedAt

	return chat, nil
}
