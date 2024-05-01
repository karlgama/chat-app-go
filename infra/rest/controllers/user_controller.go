package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	use_cases "github.com/karlgama/chat-app-go.git/application/usecases"
)

var validate *validator.Validate

type CreateUserOutput struct {
	ID        *uuid.UUID `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func CreateUser(c *gin.Context) {
	validate = validator.New(validator.WithRequiredStructEnabled())

	var input use_cases.CreateUserInput

	if bindError := c.ShouldBindJSON(&input); bindError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": bindError.Error(),
		})
		return
	}
	err := validate.Struct(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := use_cases.CreateUser(&input)

	output := CreateUserOutput{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusCreated, output)
}
