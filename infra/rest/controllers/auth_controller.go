package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	auth_use_cases "github.com/karlgama/chat-app-go.git/application/usecases/auth"
	"github.com/karlgama/chat-app-go.git/domain/entities"
	"github.com/karlgama/chat-app-go.git/infra/utils"
)

func handleLoginError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	utils.SafeLogError(c, "Login failed", map[string]interface{}{
		"error": err.Error(),
	})

	if errors.Is(err, entities.ErrInvalidCredentials) {
		utils.SafeLogWarn(c, "Login attempt with invalid credentials", map[string]interface{}{
			"reason": "invalid_credentials",
		})
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return true
	}

	if errors.Is(err, entities.ErrTokenGeneration) {
		utils.SafeLogError(c, "Token generation failed", map[string]interface{}{
			"reason": "token_generation_failed",
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return true
	}

	utils.SafeLogError(c, "Unknown login error", map[string]interface{}{
		"reason": "unknown_error",
	})
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Internal server error",
	})
	return true
}

func Login(c *gin.Context) {
	// TODO: Adicionar autenticação do google
	var input auth_use_cases.LoginInput

	if bindError := c.ShouldBindJSON(&input); bindError != nil {
		utils.SafeLogWarn(c, "Invalid request body", map[string]interface{}{
			"error": bindError.Error(),
		})

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	utils.SafeLogInfo(c, "Login attempt", map[string]interface{}{
		"email": input.Email,
	})

	token, err := auth_use_cases.Login(&input)

	if handleLoginError(c, err) {
		return
	}

	if token == nil || token.Token == nil {
		utils.SafeLogError(c, "Token is nil after successful login", map[string]interface{}{
			"reason": "nil_token",
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	utils.SafeLogInfo(c, "Login successful", map[string]interface{}{
		"email": input.Email,
	})

	output := auth_use_cases.LoginOutput{
		Token: token.Token,
	}

	c.JSON(http.StatusOK, output)
}
