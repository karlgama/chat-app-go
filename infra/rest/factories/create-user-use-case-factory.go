package factories

import (
	user_use_cases "github.com/karlgama/chat-app-go.git/application/usecases/user"
	postgre "github.com/karlgama/chat-app-go.git/infra/postgreSQL/repositories"
)

func CreateUserUseCase() *user_use_cases.CreateUserUseCase {
	return user_use_cases.NewCreateUserUseCase(&postgre.UserPostgresRepository{})
}
