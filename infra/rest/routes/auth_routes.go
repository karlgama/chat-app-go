package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/karlgama/chat-app-go.git/infra/rest/controllers"
)

func HandleAuthRequests(r *gin.Engine) {

	userGroup := r.Group("/")
	{
		userGroup.POST("/login", controllers.Login)
	}

}
