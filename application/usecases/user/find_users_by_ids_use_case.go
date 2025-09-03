package user_use_cases

import (
	"github.com/google/uuid"
	"github.com/karlgama/chat-app-go.git/application/repositories"
	"github.com/karlgama/chat-app-go.git/domain/entities"
)

type FindUsersByIdsUseCase struct {
	repositories repositories.UserRepository
}

func (f *FindUsersByIdsUseCase) FindUsersByIds(ids *[]uuid.UUID) (*[]entities.User, error) {
	users, err := f.repositories.FindUsersByIds(ids)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func NewFindUsersByIdsUseCase(repository repositories.UserRepository) *FindUsersByIdsUseCase {
	return &FindUsersByIdsUseCase{repositories: repository}
}
