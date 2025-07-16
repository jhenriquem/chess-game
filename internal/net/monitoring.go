package net

import (
	"chess-game/internal/logic"
	"chess-game/internal/model"
	"chess-game/internal/pkg/format"
	"chess-game/internal/protocol"
	"fmt"

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
				message := fmt.Sprintf("🟢 It's still your turn ( %s )", err.Error())
				protocol.SendMessage(conn, "TURN", message, true, format.ToFormatGame(game, game.Turn))
			} else {
				protocol.SendMessage(game.Turn.Client, "TURN", "🟢 It's your turn", true, format.ToFormatGame(game, game.Turn))
				for _, player := range game.Players {
					if player != game.Turn {
						protocol.SendMessage(player.Client, "WAIT", "🫸 Waiting for player to move", false, format.ToFormatGame(game, player))
					}
				}
			}
		}
	}
}
