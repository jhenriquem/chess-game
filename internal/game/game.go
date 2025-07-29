package game

import (
	"chess-game/internal/protocol"
	"chess-game/model"
	"chess-game/pkg/format"

	"github.com/gorilla/websocket"
)

func StartGame(game *model.Game) {
	// White Player message
	protocol.SendMessage(game.CurrentPlayer.Client, "INIT", "📌 You are playing, you are white ⬜. \n🟢 It's your turn  ", true, format.Game(game))

	// Black Player message
	msg := "📌 You are playing, you are black ⬛"
	protocol.SendMessage(game.GetPlayer("B").Client, "INIT", msg, false, format.Game(game))

	for {
		select {
		case player := <-game.Desconnect:
			msg := "⚠️ Outher player desconnected \n ✨ You win"

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
