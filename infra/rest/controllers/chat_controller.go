package controllers

import (
	"github.com/gin-gonic/gin"
	chat_use_cases "github.com/karlgama/chat-app-go.git/application/usecases/chat"
)

func CreateChat(c *gin.Context) {
	var input chat_use_cases.CreateChatInput
}
