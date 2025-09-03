package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	JWT      JWTConfig
	App      AppConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type JWTConfig struct {
	Secret string
}

type AppConfig struct {
	Environment string
	Port        string
	LogLevel    string
}

var AppSettings *Config

func LoadConfig() {
	// Carrega o profile do ambiente
	profile := getProfile()
	
	// Carrega o arquivo .env específico do ambiente
	envFile := fmt.Sprintf(".env.%s", profile)
	
	// Verifica se o arquivo existe
	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		log.Printf("Arquivo %s não encontrado, tentando carregar .env padrão", envFile)
		envFile = ".env"
	}
	
	// Carrega as variáveis de ambiente
	err := godotenv.Load(envFile)
	if err != nil {
		log.Printf("Aviso: Não foi possível carregar o arquivo %s: %v", envFile, err)
		log.Println("Usando variáveis de ambiente do sistema")
	} else {
		log.Printf("Configuração carregada do arquivo: %s", envFile)
	}
	
	// Inicializa a configuração
	AppSettings = &Config{
		Database: DatabaseConfig{
			Host:     getEnvOrDefault("DB_HOST", "localhost"),
			Port:     getEnvOrDefault("DB_PORT", "5432"),
			User:     getEnvOrDefault("DB_USER", "postgres"),
			Password: getEnvOrDefault("DB_PASSWORD", ""),
			Name:     getEnvOrDefault("DB_NAME", "chat_app_db"),
		},
		JWT: JWTConfig{
			Secret: getEnvOrDefault("JWT_SECRET", "default_secret_change_me"),
		},
		App: AppConfig{
			Environment: getEnvOrDefault("APP_ENV", profile),
			Port:        getEnvOrDefault("APP_PORT", "8080"),
			LogLevel:    getEnvOrDefault("LOG_LEVEL", "info"),
		},
	}
	
	log.Printf("Aplicação configurada para ambiente: %s", AppSettings.App.Environment)
}

func getProfile() string {
	// Prioridade: 1. Variável de ambiente APP_ENV, 2. Argumento da linha de comando, 3. Padrão 'local'
	if profile := os.Getenv("APP_ENV"); profile != "" {
		return profile
	}
	
	// Verifica argumentos da linha de comando
	args := os.Args
	for i, arg := range args {
		if arg == "--env" && i+1 < len(args) {
			return args[i+1]
		}
		if arg == "--profile" && i+1 < len(args) {
			return args[i+1]
		}
	}
	
	return "local"
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func GetDatabaseConnectionString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		AppSettings.Database.Host,
		AppSettings.Database.User,
		AppSettings.Database.Password,
		AppSettings.Database.Name,
		AppSettings.Database.Port)
}

func IsProduction() bool {
	return AppSettings.App.Environment == "production"
}

func IsDevelopment() bool {
	return AppSettings.App.Environment == "development"
}

func IsLocal() bool {
	return AppSettings.App.Environment == "local"
}
