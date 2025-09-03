package factories

import (
	chat_use_cases "github.com/karlgama/chat-app-go.git/application/usecases/chat"
	postgre "github.com/karlgama/chat-app-go.git/infra/postgreSQL/repositories"
)

func CreateChatUseCase() *chat_use_cases.CreateChatUseCase {
	chatRepository := &postgre.ChatPostgresRepository{}
	userRepository := &postgre.UserPostgresRepository{}
	return chat_use_cases.NewChatUseCase(chatRepository, userRepository)
}
