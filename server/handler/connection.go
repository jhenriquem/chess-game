package handler

import (
	"chess-game/model"
	"chess-game/server/game"
	"log"
	"net"
	"sync"
	"time"
)

var mutex sync.Mutex

func Connection(ln net.Listener) {
	go func() {
		var waitingConn *model.Player

		timer := time.NewTimer(2 * time.Minute)
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Printf("Erro ao conectar o client  : %s", err.Error())
			}

			player := game.NewPlayer(conn)
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
