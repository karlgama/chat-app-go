package chat_use_cases

import (
	"log"

	"github.com/google/uuid"
	"github.com/karlgama/chat-app-go.git/application/repositories"
	user_use_cases "github.com/karlgama/chat-app-go.git/application/usecases/user"
	"github.com/karlgama/chat-app-go.git/domain/entities"
)

type CreateChatInput struct {
	UserIds *[]uuid.UUID `json:"user_ids" binding:"required"`
}

type CreateChatUseCase struct {
	findUsersByIDsUseCase *user_use_cases.FindUsersByIdsUseCase
	repository            repositories.ChatRepository
}

func NewChatUseCase(repository repositories.ChatRepository, userRepository repositories.UserRepository) *CreateChatUseCase {
	findUsersByIDsUseCase := user_use_cases.NewFindUsersByIdsUseCase(userRepository)
	return &CreateChatUseCase{
		findUsersByIDsUseCase: findUsersByIDsUseCase,
		repository:            repository,
	}
}

func (c *CreateChatUseCase) CreateChat(input *CreateChatInput) (*entities.Chat, error) {
	log.Println("creating chat")
	users, err := c.findUsersByIDsUseCase.FindUsersByIds(input.UserIds)
	if err != nil {
		return nil, err
	}

	chat := &entities.Chat{
		Users: users,
	}

	return c.repository.Save(chat)

}
