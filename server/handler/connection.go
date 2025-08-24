package handler

import (
	"chess-game/model"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	mutex    sync.Mutex
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	waitingConn *model.Player
)

func Game(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Erro ao conectar o client : %s", err.Error())
	}

	go connection(conn)
}

func firstConnection(conn *websocket.Conn) *model.Player {
	var connMessage model.Message
	if err := conn.ReadJSON(&connMessage); err != nil {
		log.Println("failed to read connection:", err)
		conn.Close()
	}

	if connMessage.Type != "CONNECTED" {
		_ = conn.WriteJSON(model.Message{
			Type: "ERROR",
			Data: model.Data{Message: "First message must be CONNECTED"},
		})
		conn.Close()
	}

	conn.SetReadDeadline(time.Time{})

	player := &model.Player{
		Conn: conn,
		Name: connMessage.Data.Message,
		Send: make(chan model.Message, 10),
	}

	go func(p *model.Player) {
		defer p.Conn.Close()
		for msg := range p.Send {
			if err := p.Conn.WriteJSON(msg); err != nil {
				log.Println("erro ao enviar para", p.Name, ":", err)
				return
			}
		}
	}(player)

	return player
}

func connection(conn *websocket.Conn) {
	timer := time.NewTimer(2 * time.Minute)

	conn.SetReadDeadline(time.Now().Add(30 * time.Second))

	player := firstConnection(conn)
	if player == nil {
		return
	}

	mutex.Lock()

	if waitingConn == nil {
		player.Conn.WriteJSON(model.Message{Type: "WAIT", Data: model.Data{FEN: "", Message: "Waiting for another player."}})
		waitingConn = player
		mutex.Unlock()

		go func() {
			<-timer.C
			mutex.Lock()
			defer mutex.Unlock()
			if waitingConn == player {
				waitingConn = nil

				player.Conn.WriteJSON(model.Message{Type: "TIMEOUT", Data: model.Data{FEN: "", Message: "Waiting time is over. Please try again later."}})

				conn.Close()
			}
		}()
		return
	}

	p1 := waitingConn

	// Reset
	waitingConn = nil
	mutex.Unlock()

	go match(p1, player)
}
