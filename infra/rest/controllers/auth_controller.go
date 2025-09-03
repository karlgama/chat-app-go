package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	auth_use_cases "github.com/karlgama/chat-app-go.git/application/usecases/auth"
	"github.com/karlgama/chat-app-go.git/domain/constants"
	"github.com/karlgama/chat-app-go.git/domain/entities"
	"github.com/karlgama/chat-app-go.git/infra/utils"
)

func handleLoginError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	utils.SafeLogError(c, constants.LogLoginFailed, map[string]interface{}{
		"error": err.Error(),
	})

	if errors.Is(err, entities.ErrInvalidCredentials) {
		utils.SafeLogWarn(c, constants.LogInvalidCredentials, map[string]interface{}{
			"reason": constants.ReasonInvalidCredentials,
		})
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": constants.ErrInvalidCredentials,
		})
		return true
	}

	if errors.Is(err, entities.ErrTokenGeneration) {
		utils.SafeLogError(c, constants.LogTokenGenerationFailed, map[string]interface{}{
			"reason": constants.ReasonTokenGenerationFailed,
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": constants.ErrInternalServer,
		})
		return true
	}

	utils.SafeLogError(c, constants.LogUnknownLoginError, map[string]interface{}{
		"reason": constants.ReasonUnknownError,
	})
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": constants.ErrInternalServer,
	})
	return true
}

func Login(c *gin.Context) {
	// TODO: Adicionar autenticação do google
	var input auth_use_cases.LoginInput

	if bindError := c.ShouldBindJSON(&input); bindError != nil {
		utils.SafeLogWarn(c, constants.LogInvalidRequestBody, map[string]interface{}{
			"error": bindError.Error(),
		})

		c.JSON(http.StatusBadRequest, gin.H{
			"error": constants.ErrInvalidBody,
		})
		return
	}

	utils.SafeLogInfo(c, constants.LogLoginAttempt, map[string]interface{}{
		"email": input.Email,
	})

	token, err := auth_use_cases.Login(&input)

	if handleLoginError(c, err) {
		return
	}

	if token == nil || token.Token == nil {
		utils.SafeLogError(c, constants.LogTokenIsNil, map[string]interface{}{
			"reason": constants.ReasonNilToken,
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": constants.ErrInternalServer,
		})
		return
	}

	utils.SafeLogInfo(c, constants.LogLoginSuccessful, map[string]interface{}{
		"email": input.Email,
	})

	output := auth_use_cases.LoginOutput{
		Token: token.Token,
	}

	c.JSON(http.StatusOK, output)
}
