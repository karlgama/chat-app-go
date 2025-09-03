package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

// ConfigureLogger configura o logrus baseado no ambiente
func ConfigureLogger() {
	// Configurar nível de log baseado no ambiente
	switch AppSettings.App.LogLevel {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	// Configurar formato baseado no ambiente
	if AppSettings.App.Environment == "production" {
		// Em produção: JSON estruturado
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	} else {
		// Em desenvolvimento: formato colorido e legível
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:     true,
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}

	// Output para stdout
	logrus.SetOutput(os.Stdout)

	// Adicionar campos padrão para todos os logs
	logrus.WithFields(logrus.Fields{
		"app":     "chat-app-go",
		"version": "v0.01",
		"env":     AppSettings.App.Environment,
	}).Info("Logger configurado")
}
