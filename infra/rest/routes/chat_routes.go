package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/karlgama/chat-app-go.git/infra/rest/controllers"
)

func HandleChatRequests(r *gin.Engine) {
	chatGroup := r.Group("/")
	{
		chatGroup.POST("/chat", controllers.CreateChat)
	}
}
