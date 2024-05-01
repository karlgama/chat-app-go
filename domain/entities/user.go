package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        *uuid.UUID
	Name      string
	Email     string
	Password  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func NewUser(name string, email string, password string) *User {
	return &User{
		nil,
		name,
		email,
		password,
		nil,
		nil,
	}
}
