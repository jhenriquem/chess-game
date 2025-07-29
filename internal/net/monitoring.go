package net

import (
	"chess-game/model"

	"github.com/gorilla/websocket"
)

func MonitoringClient(conn *websocket.Conn, game *model.Game, message chan model.ClientMessage, done chan struct{}) {
	SetPongHandler(conn)
	go StartPinger(conn, done)

	ReaderClient(conn, game, message, done)

	<-done
	return
}
