package entities

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        uuid.UUID
	ChatID    uuid.UUID
	UserID    uuid.UUID
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
