package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	user_use_cases "github.com/karlgama/chat-app-go.git/application/usecases/user"
	"github.com/karlgama/chat-app-go.git/infra/rest/factories"
)

type CreateUserOutput struct {
	ID        *uuid.UUID `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func CreateUser(c *gin.Context) {
	useCase := factories.CreateUserUseCase()

	var input user_use_cases.CreateUserInput

	if bindError := c.ShouldBindJSON(&input); bindError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": bindError.Error(),
		})
		return
	}

	user, _ := useCase.CreateUser(&input)

	output := CreateUserOutput{
		ID:        user.ExternalID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusCreated, output)
}
