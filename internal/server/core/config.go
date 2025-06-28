package core

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var (
	waitingConn *websocket.Conn
	mutex       sync.Mutex
)

func Run() *string {
	http.HandleFunc("/game", handlerGame)

	return addr
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

func sendInfo(c *websocket.Conn, msg string) error {
	err := c.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		return fmt.Errorf("Error sending message : ", err)
	}
	return nil
}
