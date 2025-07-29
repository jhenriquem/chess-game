package server

import (
	"chess-game/internal/protocol"
	"chess-game/model"
	"time"

	"github.com/gorilla/websocket"
)

func HandleConnection(conn *websocket.Conn) {
	timer := time.NewTimer(1 * time.Minute)

	mutex.Lock()
	if waitingConn == nil {
		protocol.SendMessage(conn, "WAIT", "🫸 Waiting for another player.", false, model.GameFormat{})
		waitingConn = conn
		mutex.Unlock()

		<-timer.C

		mutex.Lock()
		defer mutex.Unlock()

		if waitingConn == conn {
			waitingConn = nil

			protocol.SendMessage(conn, "TIMEOUT", "❌ Waiting time is over. Please try again later.", false, model.GameFormat{})
			conn.WriteMessage(websocket.CloseMessage, []byte{})
			conn.Close()
		}

		return
	}

	p1 := waitingConn
	waitingConn = nil
	mutex.Unlock()

	msg := "🔗 Player found, starting game"

	protocol.SendMessage(p1, "PLAYER_FOUND", msg, false, model.GameFormat{})
	protocol.SendMessage(conn, "PLAYER_FOUND", msg, false, model.GameFormat{})

	go HandleMatch(p1, conn)
}
