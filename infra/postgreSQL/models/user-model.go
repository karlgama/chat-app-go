package postgreSQL

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	ID         *int
	ExternalID *uuid.UUID
	Name       string
	Email      string
	Password   string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}
