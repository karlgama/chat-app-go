package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	chat_use_cases "github.com/karlgama/chat-app-go.git/application/usecases/chat"
	"github.com/karlgama/chat-app-go.git/domain/constants"
	"github.com/karlgama/chat-app-go.git/infra/rest/factories"
)

func CreateChat(c *gin.Context) {
	var input chat_use_cases.CreateChatInput
	useCase := factories.CreateChatUseCase()

	if bindError := c.ShouldBindJSON(&input); bindError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": constants.ErrInvalidBody,
		})
		return
	}

	output, err := useCase.CreateChat(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": constants.ErrInternalServer,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         output.ID,
		"created_at": output.CreatedAt,
		"updated_at": output.UpdatedAt,
	})
}
