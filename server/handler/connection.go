package handler

import (
	"chess-game/model"
	"encoding/json"
	"log"
	"net"
	"sync"
	"time"
)

var mutex sync.Mutex

func newPlayer(conn net.Conn) *model.Player {
	enc := json.NewEncoder(conn)
	dec := json.NewDecoder(conn)

	var connMessage model.Message

	conn.SetReadDeadline(time.Now().Add(30 * time.Second))

	if err := dec.Decode(&connMessage); err != nil {
		log.Println("failed to read connection:", err)
		conn.Close()
		return nil
	}

	if connMessage.Type != "CONNECTED" {
		enc.Encode(model.Message{Type: "ERROR", Data: model.Data{FEN: "", Message: "First message must be CONNECTED"}})
		conn.Close()
		return nil
	}

	player := model.Player{
		Conn:    conn,
		Decoder: dec,
		Encoder: enc,
		Name:    connMessage.Data.Message,
	}

	return &player
}

func Connection(ln net.Listener) {
	go func() {
		var waitingConn *model.Player

		timer := time.NewTimer(2 * time.Minute)
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Printf("Erro ao conectar o client  : %s", err.Error())
			}

			player := newPlayer(conn)
			if player == nil {
				continue
			}

			mutex.Lock()
			if waitingConn == nil {
				player.Encoder.Encode(model.Message{Type: "WAIT", Data: model.Data{FEN: "", Message: "Waiting for another player."}})
				waitingConn = player
				mutex.Unlock()

				go func() {
					<-timer.C
					mutex.Lock()
					defer mutex.Unlock()
					if waitingConn == player {
						waitingConn = nil

						player.Encoder.Encode(model.Message{Type: "TIMEOUT", Data: model.Data{FEN: "", Message: "Waiting time is over. Please try again later."}})

						conn.Close()
					}
				}()

				continue
			}

			p1 := waitingConn

			// Reset
			waitingConn = nil
			mutex.Unlock()

			go Match(p1, player)
		}
	}()
}
