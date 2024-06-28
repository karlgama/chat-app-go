package postgreSQL

import (
	"github.com/google/uuid"
	"github.com/karlgama/chat-app-go.git/domain/entities"
	"github.com/karlgama/chat-app-go.git/infra/postgreSQL"
	"github.com/karlgama/chat-app-go.git/infra/postgreSQL/models"
)

type UserPostgresRepository struct{}

func (u *UserPostgresRepository) Save(user *entities.User) (*entities.User, error) {
	model := models.UserModel{
		ID:         nil,
		ExternalID: user.ExternalID,
		Name:       user.Name,
		Email:      user.Email,
		Password:   user.Password,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
	postgreSQL.DB.Create(&model)
	user.ID = model.ID
	user.CreatedAt = model.CreatedAt
	user.UpdatedAt = model.UpdatedAt

	return user, nil
}

func (u *UserPostgresRepository) FindUsersByIds(ids *[]uuid.UUID) (*[]entities.User, error) {

}
