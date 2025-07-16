package game

import (
	"chess-game/internal/model"
	"chess-game/internal/net"
	"chess-game/internal/pkg/format"
	"chess-game/internal/protocol"
)

func Run(game *model.Game) {
	protocol.SendMessage(game.Turn.Client, "INIT", "📌 You are playing, you are white ⬜. \n🟢 It's your turn  ", true, format.ToFormatGame(game, game.Turn))
	for _, p := range game.Players {
		if p.Color != game.Turn.Color {
			msg := "📌 You are playing, you are black ⬛"

			protocol.SendMessage(p.Client, "INIT", msg, false, format.ToFormatGame(game, p))
		}
	}

	for {
		select {
		case player := <-game.Desconnect:
			msg := "⚠️ Outher player desconnected \n ✨ You win"
			if player.Color == "black" {
				outherPlayer := GetPlayer(game, "white")
				protocol.SendMessage(outherPlayer.Client, "DESCONNECTED", msg, false, format.ToFormatGame(game, outherPlayer))
				net.CloseConnection(outherPlayer.Client)
			} else {
				outherPlayer := GetPlayer(game, "black")
				protocol.SendMessage(outherPlayer.Client, "DESCONNECTED", msg, false, format.ToFormatGame(game, outherPlayer))
				net.CloseConnection(outherPlayer.Client)
			}

			break
		}
	}
}
