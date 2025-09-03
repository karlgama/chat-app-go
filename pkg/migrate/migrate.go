package migrate

import (
	"fmt"

	"github.com/karlgama/chat-app-go.git/infra/postgreSQL"
)

// RunMigrations executa todas as migrações pendentes
func RunMigrations() error {
	// Configura banco de dados
	postgreSQL.SetupDatabase()

	// Executa auto-migração
	postgreSQL.AutoMigrate()

	// Executa migrações customizadas
	postgreSQL.RunCustomMigrations()

	return nil
}

// DropAllTables remove todas as tabelas do banco
func DropAllTables() error {
	// Configura banco se não estiver configurado
	postgreSQL.SetupDatabase()

	// Acessa a variável DB diretamente do package postgreSQL
	// Remove todas as tabelas usando CASCADE
	if err := postgreSQL.DB.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;").Error; err != nil {
		return fmt.Errorf("erro ao remover tabelas: %w", err)
	}

	return nil
}

// ResetDatabase reseta completamente o banco (drop + migrate)
func ResetDatabase() error {
	if err := DropAllTables(); err != nil {
		return fmt.Errorf("erro ao remover tabelas: %w", err)
	}

	if err := RunMigrations(); err != nil {
		return fmt.Errorf("erro ao executar migrações: %w", err)
	}

	return nil
}
