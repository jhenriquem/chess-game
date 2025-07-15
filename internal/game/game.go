package game

import (
	"chess-game/internal/model"
	"chess-game/internal/net"
	"chess-game/internal/pkg/format"
	"chess-game/internal/protocol"
	"fmt"
)

func Run(game *model.Game) {
	for _, p := range game.Players {
		msg := fmt.Sprintf("You are playing, you are %s", p.Color)
		protocol.SendMessage(p.Client, "initGame", msg, format.ToFormatGame(game))
	}

	for {
		select {
		case player := <-game.Desconnect:
			msg := "Outher player desconnected, you win"
			if player.Color == "black" {
				protocol.SendMessage(GetPlayer(game, "white").Client, "desconnected", msg, format.ToFormatGame(game))
				net.CloseConnection(GetPlayer(game, "white").Client)
			} else {
				protocol.SendMessage(GetPlayer(game, "black").Client, "desconnected", msg, format.ToFormatGame(game))
				net.CloseConnection(GetPlayer(game, "black").Client)
			}

			break
		}
	}
}
