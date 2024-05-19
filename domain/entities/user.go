package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         *int
	ExternalID *uuid.UUID
	Name       string
	Email      string
	Password   string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

func NewUser(name string, email string, password string) *User {
	id, _ := uuid.NewRandom()
	return &User{
		nil,
		&id,
		name,
		email,
		password,
		nil,
		nil,
	}
}
