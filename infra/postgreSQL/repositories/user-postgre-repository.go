package postgreSQL

import (
	"errors"

	"github.com/google/uuid"
	"github.com/karlgama/chat-app-go.git/domain/entities"
	"github.com/karlgama/chat-app-go.git/infra/postgreSQL"
	"github.com/karlgama/chat-app-go.git/infra/postgreSQL/models"
	"gorm.io/gorm"
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
	result := postgreSQL.DB.Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}

	user.ID = model.ID
	user.CreatedAt = model.CreatedAt
	user.UpdatedAt = model.UpdatedAt

	return user, nil
}

func (u *UserPostgresRepository) FindUsersByIds(ids *[]uuid.UUID) (*[]entities.User, error) {
	if ids == nil || len(*ids) == 0 {
		return &[]entities.User{}, nil
	}

	var models []models.UserModel
	result := postgreSQL.DB.Where("externalID IN ?", *ids).Find(&models)

	if result.Error != nil {
		return nil, result.Error
	}

	users := make([]entities.User, len(models))
	for i, model := range models {
		users[i] = entities.User{
			ID:         model.ID,
			ExternalID: model.ExternalID,
			Name:       model.Name,
			Email:      model.Email,
			Password:   model.Password,
			CreatedAt:  model.CreatedAt,
			UpdatedAt:  model.UpdatedAt,
		}
	}

	return &users, nil
}

func (u *UserPostgresRepository) FindUserByEmail(email string) (*entities.User, error) {
	var model models.UserModel
	result := postgreSQL.DB.Where("email = ?", email).First(&model)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}

	user := &entities.User{
		ID:         model.ID,
		ExternalID: model.ExternalID,
		Name:       model.Name,
		Email:      model.Email,
		Password:   model.Password,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
	}

	return user, nil
}
