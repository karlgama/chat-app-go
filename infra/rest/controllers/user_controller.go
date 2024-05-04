package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	use_cases "github.com/karlgama/chat-app-go.git/application/usecases/user"
)

type CreateUserOutput struct {
	ID        *uuid.UUID `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func CreateUser(c *gin.Context) {

	var input use_cases.CreateUserInput

	if bindError := c.ShouldBindJSON(&input); bindError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": bindError.Error(),
		})
		return
	}

	user, _ := use_cases.CreateUser(&input)

	output := CreateUserOutput{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusCreated, output)
}
