package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var (
	waitingConn *websocket.Conn
	mutex       sync.Mutex
)

func Run() *string {
	addr := fmt.Sprintf("localhost:%s", os.Getenv("PORT"))

	http.HandleFunc("/game", handlerGame)

	return &addr
}

func handlerGame(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	log.Println("Connected client")

	go HandleConnection(c)
}
