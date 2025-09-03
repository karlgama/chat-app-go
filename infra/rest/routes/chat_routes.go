package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/karlgama/chat-app-go.git/infra/rest/controllers"
)

func HandleChatRequests(r *gin.RouterGroup) {
	chatGroup := r.Group("/chats")
	{
		chatGroup.POST("", controllers.CreateChat)
	}
}
