package postgreSQL

import (
	"log"

	"github.com/karlgama/chat-app-go.git/infra/postgreSQL/models"
)

// RunCustomMigrations executa migrações customizadas que o AutoMigrate não consegue fazer
func RunCustomMigrations() {
	log.Println("🔧 Executando migrações customizadas...")

	// Adiciona índices customizados se necessário
	createCustomIndexes()

	// Adiciona constraints customizadas
	createCustomConstraints()

	// Adiciona dados iniciais se necessário
	seedInitialData()

	log.Println("✅ Migrações customizadas concluídas!")
}

func createCustomIndexes() {
	// Exemplo: Índice composto para performance
	err := DB.Exec(`
		CREATE INDEX IF NOT EXISTS idx_chat_users_chat_user 
		ON chat_users(chat_id, user_id)
	`).Error

	if err != nil {
		log.Printf("⚠️  Erro ao criar índice customizado: %v", err)
	}
}

func createCustomConstraints() {
	// Verifica se a constraint já existe antes de criar
	var count int64

	// Foreign key para chat_id
	DB.Raw(`
		SELECT COUNT(*) 
		FROM information_schema.table_constraints 
		WHERE constraint_name = 'fk_chat_users_chat_id' 
		AND table_name = 'chat_users'
	`).Scan(&count)

	if count == 0 {
		err := DB.Exec(`
			ALTER TABLE chat_users 
			ADD CONSTRAINT fk_chat_users_chat_id 
			FOREIGN KEY (chat_id) REFERENCES chats(id) ON DELETE CASCADE
		`).Error

		if err != nil {
			log.Printf("⚠️  Erro ao criar constraint chat_id: %v", err)
		} else {
			log.Println("✅ Constraint fk_chat_users_chat_id criada")
		}
	} else {
		log.Println("✅ Constraint fk_chat_users_chat_id já existe")
	}

	// Foreign key para user_id
	DB.Raw(`
		SELECT COUNT(*) 
		FROM information_schema.table_constraints 
		WHERE constraint_name = 'fk_chat_users_user_id' 
		AND table_name = 'chat_users'
	`).Scan(&count)

	if count == 0 {
		err := DB.Exec(`
			ALTER TABLE chat_users 
			ADD CONSTRAINT fk_chat_users_user_id 
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		`).Error

		if err != nil {
			log.Printf("⚠️  Erro ao criar constraint user_id: %v", err)
		} else {
			log.Println("✅ Constraint fk_chat_users_user_id criada")
		}
	} else {
		log.Println("✅ Constraint fk_chat_users_user_id já existe")
	}
}

func seedInitialData() {
	// Verifica se já existem dados
	var userCount int64
	DB.Model(&models.UserModel{}).Count(&userCount)

	if userCount == 0 {
		log.Println("📋 Criando dados iniciais de exemplo...")
		// Aqui você pode criar usuários de exemplo para desenvolvimento
		// Apenas se necessário
	}
}

// DropAllTables remove todas as tabelas (CUIDADO! Apenas para desenvolvimento)
func DropAllTables() {
	if DB == nil {
		log.Println("❌ Banco de dados não conectado")
		return
	}

	log.Println("🗑️  ATENÇÃO: Removendo todas as tabelas...")

	// Lista das tabelas para remover na ordem correta (devido às foreign keys)
	tables := []string{
		"chat_users",
		"chats",
		"users",
	}

	for _, table := range tables {
		err := DB.Exec("DROP TABLE IF EXISTS " + table + " CASCADE").Error
		if err != nil {
			log.Printf("⚠️  Erro ao remover tabela %s: %v", table, err)
		} else {
			log.Printf("✅ Tabela %s removida", table)
		}
	}

	log.Println("🗑️  Todas as tabelas foram removidas!")
}
