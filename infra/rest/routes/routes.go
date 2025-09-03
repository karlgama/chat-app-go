package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karlgama/chat-app-go.git/infra/config"
	"github.com/karlgama/chat-app-go.git/infra/rest/middlewares"
	security_middlewares "github.com/karlgama/chat-app-go.git/infra/security/middlewares"
)

func SetupRoutes() {
	// Configura o modo do Gin baseado no ambiente
	if config.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()

	// Adiciona middleware de trace ID para todas as rotas
	r.Use(middlewares.TraceIDMiddleware())

	HandleUserRequests(r)
	HandleAuthRequests(r)

	protectedRoutes := r.Group("/")
	protectedRoutes.Use(security_middlewares.AuthenticationMiddleware())
	{
		protectedRoutes.GET("/protected", func(r *gin.Context) {
			r.JSON(http.StatusOK, gin.H{"message": "protected route"})
		})
		HandleChatRequests(protectedRoutes)
	}

	port := ":" + config.AppSettings.App.Port
	r.Run(port)
}
