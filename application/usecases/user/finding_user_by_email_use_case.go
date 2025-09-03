package user_use_cases

import (
	"github.com/karlgama/chat-app-go.git/application/repositories"
	"github.com/karlgama/chat-app-go.git/domain/entities"
	postgre "github.com/karlgama/chat-app-go.git/infra/postgreSQL/repositories"
)

func FindUserByEmail(email string) *entities.User {
	repository := &postgre.UserPostgresRepository{}
	user, err := repository.FindUserByEmail(email)
	
	if err != nil {
		return nil
	}

	return user
}

type FindUserByEmailUseCase struct {
	repository repositories.UserRepository
}

func (f *FindUserByEmailUseCase) Execute(email string) (*entities.User, error) {
	return f.repository.FindUserByEmail(email)
}

func NewFindUserByEmailUseCase(repository repositories.UserRepository) *FindUserByEmailUseCase {
	return &FindUserByEmailUseCase{repository: repository}
}
