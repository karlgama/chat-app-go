package main

import (
	"fmt"
	"log"
	"os"

	"github.com/karlgama/chat-app-go.git/infra/config"
	"github.com/karlgama/chat-app-go.git/infra/postgreSQL"
	"github.com/karlgama/chat-app-go.git/infra/rest/routes"
	"github.com/karlgama/chat-app-go.git/pkg/migrate"
)

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,

// 	CheckOrigin: func(r *http.Request) bool { return true },
// }

// func reader(conn *websocket.Conn) {
// 	for {
// 		messageType, p, err := conn.ReadMessage()

// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}

// 		fmt.Println(string(p))

// 		if err := conn.WriteMessage(messageType, p); err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}
// }

// func serveWs(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println(r.Host)

// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	reader(ws)
// }

// func setupRoutesWS() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Simple server")
// 	})
// 	http.HandleFunc("/ws", wsEndpoint)
// }

// func wsEndpoint(w http.ResponseWriter, r *http.Request) {
// 	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	log.Println("Client Connected")
// 	err = ws.WriteMessage(1, []byte("Hi Client!"))
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	reader(ws)
// }

func handleMigrationCommands() bool {
	if len(os.Args) <= 1 {
		return false
	}

	command := os.Args[1]

	switch command {
	case "migrate":
		fmt.Println("Executando migrações...")
		if err := migrate.RunMigrations(); err != nil {
			fmt.Printf("Erro ao executar migrações: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Migrações executadas com sucesso!")
		return true
	case "drop":
		fmt.Println("Removendo todas as tabelas...")
		if err := migrate.DropAllTables(); err != nil {
			fmt.Printf("Erro ao remover tabelas: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Tabelas removidas com sucesso!")
		return true
	case "reset":
		fmt.Println("Resetando banco de dados...")
		if err := migrate.ResetDatabase(); err != nil {
			fmt.Printf("Erro ao resetar banco de dados: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Banco de dados resetado com sucesso!")
		return true
	case "--help", "-h":
		printHelp()
		return true
	default:
		fmt.Printf("Comando desconhecido: %s\n", command)
		printHelp()
		return true
	}
}

func main() {
	config.LoadConfig()
	config.ConfigureLogger()

	if handleMigrationCommands() {
		return
	}

	if config.IsDevelopment() || config.IsLocal() {
		postgreSQL.CreateDatabaseIfNotExists()
	}

	postgreSQL.SetupDatabase()

	if config.IsDevelopment() || config.IsLocal() {
		postgreSQL.AutoMigrate()
		postgreSQL.RunCustomMigrations()
	}

	log.Printf("Chat App v0.01 - Ambiente: %s", config.AppSettings.App.Environment)
	fmt.Printf("Servidor iniciando na porta %s\n", config.AppSettings.App.Port)

	routes.SetupRoutes()
}

func printHelp() {
	fmt.Println("Chat App - Sistema de Migrações")
	fmt.Println("")
	fmt.Println("Uso:")
	fmt.Println("  go run cmd/main.go [comando]")
	fmt.Println("")
	fmt.Println("Comandos disponíveis:")
	fmt.Println("  migrate    Executa todas as migrações pendentes")
	fmt.Println("  drop       Remove todas as tabelas do banco")
	fmt.Println("  reset      Reseta o banco (drop + migrate)")
	fmt.Println("  --help     Mostra esta ajuda")
	fmt.Println("")
	fmt.Println("Exemplos:")
	fmt.Println("  go run cmd/main.go migrate")
	fmt.Println("  go run cmd/main.go reset")
	fmt.Println("")
	fmt.Println("Sem argumentos: inicia o servidor web")
}
