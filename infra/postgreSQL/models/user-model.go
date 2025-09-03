package models

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	ID         *int       `gorm:"column:id;primaryKey;autoIncrement"`
	ExternalID *uuid.UUID `gorm:"column:externalid;type:uuid;not null"`
	Name       string     `gorm:"column:name;size:255;not null"`
	Email      string     `gorm:"column:email;size:255;not null;uniqueIndex"`
	Password   string     `gorm:"column:password;size:255;not null"`
	CreatedAt  *time.Time `gorm:"column:createdat"`
	UpdatedAt  *time.Time `gorm:"column:updatedat"`
}

// TableName especifica o nome da tabela no banco
func (UserModel) TableName() string {
	return "users"
}
