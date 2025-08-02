package handler

import (
	"chess-game/internal/logic"
	"chess-game/internal/protocol"
	"chess-game/model"
	"chess-game/pkg/format"

	"github.com/gorilla/websocket"
)

// Timeout or Desconnected
func HandleClientExit(game *model.Game) {
	for {
		select {
		case <-game.Timeout:
			msg1 := "⚠️ The other player's time is up\n✨ You win"
			msg2 := "⚠️ Your time is up "

			// Close clock
			logic.StopClock(game.CurrentPlayer)

			// Close connection
			protocol.SendMessage(game.CurrentPlayer.Client, "TIMEOUT", msg2, false, format.Game(game))
			game.CurrentPlayer.Client.WriteMessage(websocket.CloseMessage, []byte{})
			game.CurrentPlayer.Client.Close()

			protocol.SendMessage(game.GetAnoutherPlayer().Client, "TIMEOUT", msg1, false, format.Game(game))
			game.GetAnoutherPlayer().Client.WriteMessage(websocket.CloseMessage, []byte{})
			game.GetAnoutherPlayer().Client.Close()

			return

		case player := <-game.Desconnect:
			msg := "⚠️ Outher player desconnected \n ✨ You win"

			// Close clock
			logic.StopClock(game.CurrentPlayer)

			for _, p := range game.Players {
				if p.Client != player.Client {
					protocol.SendMessage(p.Client, "DESCONNECTED", msg, false, format.Game(game))
					p.Client.WriteMessage(websocket.CloseMessage, []byte{})
					p.Client.Close()

				}
			}

			return
		}
	}
}
