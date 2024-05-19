package postgreSQL

import (
	"github.com/karlgama/chat-app-go.git/domain/entities"
	"github.com/karlgama/chat-app-go.git/infra/postgreSQL"
)

type UserPostgreRepository struct{}

func (u *UserPostgreRepository) Save(user *entities.User) (*entities.User, error) {
	postgreSQL.
		model := postgreSQL.UserModel{
		ExternalID: user.ExternalID,
		Name:       user.Name,
		Email:      user.Email,
		Password:   user.Password,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
	postgreSQL.DB.Create(&user)
}
