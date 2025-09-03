package models

import (
	"time"

	"github.com/google/uuid"
)

type ChatModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName especifica o nome da tabela no banco
func (ChatModel) TableName() string {
	return "chats"
}

type ChatUserModel struct {
	ChatID uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID int       `gorm:"primaryKey"`
}

// TableName especifica o nome da tabela no banco
func (ChatUserModel) TableName() string {
	return "chat_users"
}
