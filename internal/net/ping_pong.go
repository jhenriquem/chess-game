package net

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	pongWait   = 30 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

func SetPongHandler(conn *websocket.Conn) {
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
}

func SetPingHandler(conn *websocket.Conn) {
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPingHandler(func(appData string) error {
		conn.SetReadDeadline(time.Now().Add(pongWait))
		return conn.WriteControl(websocket.PongMessage, nil, time.Now().Add(10*time.Second))
	})
}

func StartPinger(conn *websocket.Conn, done <-chan struct{}) {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			if err := conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(pongWait)); err != nil {
				log.Println("Erro ao enviar ping:", err)
				return
			}
		}
	}
}
