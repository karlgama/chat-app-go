package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/karlgama/chat-app-go.git/infra/rest/controllers"
)

func HandleUserRequest(r *gin.Engine) {

	userGroup := r.Group("/users")
	{
		userGroup.POST("", controllers.CreateUser)
	}

}
