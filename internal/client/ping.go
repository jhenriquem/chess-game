package client

import (
	"time"

	"github.com/gorilla/websocket"
)

func SetPinglogic(conn *websocket.Conn) {
	conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	conn.SetPingHandler(func(appData string) error {
		conn.SetReadDeadline(time.Now().Add(30 * time.Second))
		conn.WriteControl(websocket.PongMessage, nil, time.Now().Add(30*time.Second))
		return nil
	})
}
