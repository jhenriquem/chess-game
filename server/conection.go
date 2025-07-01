package server

import (
	"time"

	"github.com/gorilla/websocket"
)

func HandleConnection(conn *websocket.Conn) {
	timer := time.NewTimer(1 * time.Minute)

	waitCh := make(chan *websocket.Conn, 1)

	go func() {
		mutex.Lock()
		if waitingConn == nil {
			sendInfo(conn, "Waiting for another player.")
			waitingConn = conn
			mutex.Unlock()

			<-timer.C

			mutex.Lock()
			if waitingConn == conn {
				waitingConn = nil
				mutex.Unlock()

				sendInfo(conn, "Waiting time is over. Please try again later.")
				conn.Close()
			} else {
				mutex.Unlock()
			}

			return
		}

		p1 := waitingConn
		waitingConn = nil
		mutex.Unlock()
		waitCh <- p1
	}()

	select {
	case p1 := <-waitCh:
		msg := "Player found, starting game"

		waitingConn = nil

		sendInfo(p1, msg)
		sendInfo(conn, msg)

		go HandleMatch(p1, conn)

	case <-timer.C:
		sendInfo(conn, "Error starting game")
		conn.Close()
	}
}
