package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	chat_use_cases "github.com/karlgama/chat-app-go.git/application/usecases/chat"
	"github.com/karlgama/chat-app-go.git/infra/rest/factories"
)

func CreateChat(c *gin.Context) {
	var input chat_use_cases.CreateChatInput
	useCase := factories.CreateChatUseCase()

	if bindError := c.ShouldBindJSON(&input); bindError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": bindError.Error(),
		})
		return
	}

	output, err := useCase.CreateChat(input)

}
