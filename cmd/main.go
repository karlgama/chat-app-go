package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/karlgama/chat-app-go.git/infra/postgreSQL"
	"github.com/karlgama/chat-app-go.git/infra/rest/routes"
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

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadEnv()
	postgreSQL.SetupDatabase()
	fmt.Println("Chat App v0.01")
	routes.SetupRoutes()
}
