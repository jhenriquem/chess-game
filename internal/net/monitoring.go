package net

import (
	"chess-game/internal/logic"
	"chess-game/internal/protocol"
	"chess-game/model"
	"chess-game/pkg/format"
	"fmt"

	"github.com/gorilla/websocket"
)

func MonitoringClient(conn *websocket.Conn, game *model.Game, done chan struct{}) {
	SetPongHandler(conn)
	go StartPinger(conn, done)

	message := make(chan model.ClientMessage)

	go ReaderClient(conn, game, message, done)

	for {
		select {
		case msg := <-message:
			if result, err := logic.MovesLogic(msg, game); err != nil {
				message := fmt.Sprintf("🟢 It's still your turn ( %s )", err.Error())
				protocol.SendMessage(conn, "TURN", message, true, format.ToFormatGame(game, game.Turn))
			} else {
				game.MoveResult = result

				protocol.SendMessage(game.Turn.Client, "TURN", "🟢 It's your turn", true, format.ToFormatGame(game, game.Turn))
				for _, player := range game.Players {
					if player != game.Turn {
						protocol.SendMessage(player.Client, "WAIT", "🫸 Waiting for player to move", false, format.ToFormatGame(game, player))
					}
				}
			}
		case <-done:
			return
		}
	}
}
