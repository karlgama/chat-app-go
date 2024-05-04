package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	security_middlewares "github.com/karlgama/chat-app-go.git/infra/security/middlewares"
)

func SetupRoutes() {

	r := gin.Default()
	HandleUserRequests(r)
	HandleAuthRequests(r)

	protectedRoutes := r.Group("/")
	protectedRoutes.Use(security_middlewares.AuthenticationMiddleware())
	{
		protectedRoutes.GET("/protected", func(r *gin.Context) {
			r.JSON(http.StatusOK, gin.H{"message": "protected route"})
		})
	}

	r.Run(":8080")
}
