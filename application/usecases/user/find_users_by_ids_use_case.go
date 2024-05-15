package user_use_cases

import "github.com/karlgama/chat-app-go.git/domain/entities"

type FindUsersByIdsUseCase struct {
}

func FindUsersByIds(ids []string) []*entities.User {

	return make([]*entities.User, 0)
}
