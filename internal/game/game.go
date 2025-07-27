package game

import (
	"chess-game/internal/net"
	"chess-game/internal/protocol"
	"chess-game/model"
	"chess-game/pkg/format"
	"fmt"
)

func Run(game *model.Game) {
	fmt.Println("Game RUN")
	// White Player message
	protocol.SendMessage(game.Turn.Client, "INIT", "📌 You are playing, you are white ⬜. \n🟢 It's your turn  ", true, format.ToFormatGame(game, game.Turn))

	// Black Player message
	msg := "📌 You are playing, you are black ⬛"
	protocol.SendMessage(game.GetPlayer("B").Client, "INIT", msg, false, format.ToFormatGame(game, game.GetPlayer("B")))

	for {
		select {
		case player := <-game.Desconnect:
			msg := "⚠️ Outher player desconnected \n ✨ You win"

			for _, p := range game.Players {
				if p.Client != player.Client {
					protocol.SendMessage(p.Client, "DESCONNECTED", msg, false, format.ToFormatGame(game, p))
					net.CloseConnection(p.Client)
				}
			}

			return
		}
	}
}
