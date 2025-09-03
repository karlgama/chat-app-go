package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/karlgama/chat-app-go.git/infra/config"
	"github.com/karlgama/chat-app-go.git/infra/rest/middlewares"
	"github.com/sirupsen/logrus"
)

// LogLevel define os níveis de sensibilidade dos dados
type LogLevel string

const (
	// LogLevelPublic dados que podem ser logados em qualquer ambiente
	LogLevelPublic LogLevel = "public"
	// LogLevelSensitive dados sensíveis que só devem ser logados em debug local
	LogLevelSensitive LogLevel = "sensitive"
)

// SafeLogFields cria um mapa de campos seguros para logging baseado no ambiente
func SafeLogFields(c *gin.Context, fields map[string]interface{}) logrus.Fields {
	safeFields := logrus.Fields{
		"trace_id": middlewares.GetTraceID(c),
		"ip":       c.ClientIP(),
		"endpoint": c.Request.Method + " " + c.Request.URL.Path,
	}

	// Determina se pode logar dados sensíveis
	canLogSensitive := canLogSensitiveData()

	for key, value := range fields {
		if isSensitiveField(key) {
			if canLogSensitive {
				safeFields[key] = value
			} else {
				safeFields[key+"_masked"] = maskSensitiveData(value)
			}
		} else {
			safeFields[key] = value
		}
	}

	return safeFields
}

// SafeLogWithFields faz log estruturado seguro
func SafeLogWithFields(c *gin.Context, level logrus.Level, message string, fields map[string]interface{}) {
	safeFields := SafeLogFields(c, fields)
	logrus.WithFields(safeFields).Log(level, message)
}

// SafeLogInfo log de nível info com campos seguros
func SafeLogInfo(c *gin.Context, message string, fields map[string]interface{}) {
	SafeLogWithFields(c, logrus.InfoLevel, message, fields)
}

// SafeLogWarn log de nível warn com campos seguros
func SafeLogWarn(c *gin.Context, message string, fields map[string]interface{}) {
	SafeLogWithFields(c, logrus.WarnLevel, message, fields)
}

// SafeLogError log de nível error com campos seguros
func SafeLogError(c *gin.Context, message string, fields map[string]interface{}) {
	SafeLogWithFields(c, logrus.ErrorLevel, message, fields)
}

// canLogSensitiveData determina se dados sensíveis podem ser logados
func canLogSensitiveData() bool {
	// Só permite dados sensíveis em ambiente local com nível debug
	return config.IsLocal() && config.AppSettings.App.LogLevel == "debug"
}

// isSensitiveField verifica se um campo contém dados sensíveis
func isSensitiveField(fieldName string) bool {
	sensitiveFields := []string{
		"email", "password", "token", "jwt", "secret",
		"phone", "cpf", "cnpj", "credit_card", "ssn",
		"name", "full_name", "address", "user_id",
	}

	fieldLower := strings.ToLower(fieldName)
	for _, sensitive := range sensitiveFields {
		if strings.Contains(fieldLower, sensitive) {
			return true
		}
	}
	return false
}

// maskSensitiveData mascara dados sensíveis
func maskSensitiveData(data interface{}) string {
	str := ""
	switch v := data.(type) {
	case string:
		str = v
	default:
		return "[MASKED]"
	}

	if len(str) <= 3 {
		return "[MASKED]"
	}

	// Para emails, mostra apenas o primeiro caractere e o domínio
	if strings.Contains(str, "@") {
		parts := strings.Split(str, "@")
		if len(parts) == 2 {
			return string(str[0]) + "***@" + parts[1]
		}
	}

	// Para outros dados, mostra apenas os primeiros 2 caracteres
	return str[:2] + strings.Repeat("*", len(str)-2)
}
