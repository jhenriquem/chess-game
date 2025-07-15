package net

import (
	"chess-game/internal/logic"
	"chess-game/internal/model"
	"chess-game/internal/pkg/format"
	"chess-game/internal/protocol"

	"github.com/gorilla/websocket"
)

func MonitoringClient(conn *websocket.Conn, game *model.Game) chan struct{} {
	done := make(chan struct{})

	defer func() {
		close(done)
		return
	}()

	SetPongHandler(conn)
	go StartPinger(conn, done)

	message := make(chan model.ClientMessage)

	go ReaderClient(conn, game, message)

	for {
		select {
		case msg := <-message:
			if err := logic.MovesLogic(msg, game); err != nil {
				protocol.SendMessage(conn, "playerMove", err.Error(), format.ToFormatGame(game))
			} else {
				for _, conn := range game.Players {
					protocol.SendMessage(conn.Client, "playerMove", "Update", format.ToFormatGame(game))
				}
			}
		}
	}
}
