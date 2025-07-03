package server

import (
	"chess-game/pkg/protocol"
	"time"

	"github.com/gorilla/websocket"
)

func HandleConnection(conn *websocket.Conn) {
	timer := time.NewTimer(1 * time.Minute)

	mutex.Lock()
	if waitingConn == nil {
		protocol.SendMessage(conn, "Waiting for another player.", protocol.Game{})
		waitingConn = conn
		mutex.Unlock()

		<-timer.C

		mutex.Lock()
		defer mutex.Unlock()

		if waitingConn == conn {
			waitingConn = nil

			protocol.SendMessage(conn, "Waiting time is over. Please try again later.", protocol.Game{})
			conn.WriteMessage(websocket.CloseMessage, []byte{})
			conn.Close()
		}

		return
	}

	p1 := waitingConn
	waitingConn = nil
	mutex.Unlock()

	msg := "Player found, starting game"

	protocol.SendMessage(p1, msg, protocol.Game{})
	protocol.SendMessage(conn, msg, protocol.Game{})

	go HandleMatch(p1, conn)
}
