package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	auth_use_cases "github.com/karlgama/chat-app-go.git/application/usecases/auth"
)

func Login(c *gin.Context) {
	// TODO: Adicionar autenticação do google
	var input auth_use_cases.LoginInput

	if bindError := c.ShouldBindJSON(&input); bindError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": bindError.Error(),
		})
		return
	}

	token, _ := auth_use_cases.Login(&input)

	output := auth_use_cases.LoginOutput{
		Token: token.Token,
	}

	c.JSON(http.StatusOK, output)
}
