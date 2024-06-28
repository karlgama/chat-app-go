package factories

func CreateChatUseCase() *chat_use_cases.CreateChatUseCase {
	return chat_use_cases.NewChatUseCase(&postgre.ChatPostgresRepository{})
}
