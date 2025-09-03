package postgreSQL

import (
	"fmt"
	"log"
	"strings"

	"github.com/karlgama/chat-app-go.git/infra/config"
	"github.com/karlgama/chat-app-go.git/infra/postgreSQL/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

func SetupDatabase() {
	connString := config.GetDatabaseConnectionString()

	// Configura o logger baseado no ambiente
	var logLevel logger.LogLevel
	switch config.AppSettings.App.LogLevel {
	case "debug":
		logLevel = logger.Info
	case "info":
		logLevel = logger.Warn
	default:
		logLevel = logger.Error
	}

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	}

	DB, err = gorm.Open(postgres.Open(connString), gormConfig)
	if err != nil {
		log.Panic("failed to connect database", err)
	}

	log.Printf("Conectado ao banco de dados: %s@%s:%s/%s",
		config.AppSettings.Database.User,
		config.AppSettings.Database.Host,
		config.AppSettings.Database.Port,
		config.AppSettings.Database.Name)
}

func AutoMigrate() {
	log.Println("🔄 Iniciando auto-migração das tabelas...")

	// Lista todos os modelos que devem ser migrados
	modelsToMigrate := []interface{}{
		&models.UserModel{},
		&models.ChatModel{},
		&models.ChatUserModel{},
	}

	// Executa a migração
	err := DB.AutoMigrate(modelsToMigrate...)

	if err != nil {
		log.Printf("❌ Erro durante auto-migração: %v", err)
		log.Panic("Falha na migração do banco de dados")
	} else {
		log.Println("✅ Auto-migração concluída com sucesso!")
		log.Printf("📊 Tabelas migradas: users, chats, chat_users")
	}
}

// CreateDatabaseIfNotExists tenta criar o banco de dados se ele não existir
func CreateDatabaseIfNotExists() {
	log.Println("🔍 Verificando se o banco de dados existe...")

	// Conecta ao postgres sem especificar database para criar se necessário
	tempConfig := config.AppSettings.Database
	tempConnString := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable",
		tempConfig.Host, tempConfig.User, tempConfig.Password, tempConfig.Port)

	tempDB, err := gorm.Open(postgres.Open(tempConnString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Printf("⚠️  Não foi possível conectar ao PostgreSQL: %v", err)
		return
	}

	// Tenta criar o banco de dados
	sqlDB, _ := tempDB.DB()
	defer sqlDB.Close()

	createDBSQL := fmt.Sprintf("CREATE DATABASE %s", tempConfig.Name)
	result := tempDB.Exec(createDBSQL)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "already exists") {
			log.Printf("✅ Banco de dados '%s' já existe", tempConfig.Name)
		} else {
			log.Printf("⚠️  Erro ao criar banco de dados: %v", result.Error)
		}
	} else {
		log.Printf("✅ Banco de dados '%s' criado com sucesso!", tempConfig.Name)
	}
}
