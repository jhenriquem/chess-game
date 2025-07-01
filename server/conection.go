package server

import (
	"time"

	"github.com/gorilla/websocket"
)

func HandleConnection(conn *websocket.Conn) {
	timer := time.NewTimer(1 * time.Minute)

	mutex.Lock()
	if waitingConn == nil {
		sendInfo(conn, "Waiting for another player.")
		waitingConn = conn
		mutex.Unlock()

		<-timer.C

		mutex.Lock()
		defer mutex.Unlock()

		if waitingConn == conn {
			waitingConn = nil

			sendInfo(conn, "Waiting time is over. Please try again later.")
			conn.WriteMessage(websocket.CloseMessage, []byte{})
			conn.Close()
		}

		return
	}

	p1 := waitingConn
	waitingConn = nil
	mutex.Unlock()

	msg := "Player found, starting game"

	sendInfo(p1, msg)
	sendInfo(conn, msg)

	go HandleMatch(p1, conn)
}
