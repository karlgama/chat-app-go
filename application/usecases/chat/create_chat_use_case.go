package chat_use_cases

import (
	"github.com/google/uuid"
	user_use_cases "github.com/karlgama/chat-app-go.git/application/usecases/user"
)

type CreateChatInput struct {
	userIds *[]uuid.UUID
}

type CreateChatUseCase struct {
	findUsersByIDsUseCase *user_use_cases.FindUsersByIdsUseCase
}

func CreateChat(input *CreateChatInput) {

}
