package entities

import (
	"time"

	"github.com/google/uuid"
)

type chat struct {
	ID        uuid.UUID
	Users     []User
	Messages  []Message
	CreatedAt time.Time
	UpdatedAt time.Time
}
