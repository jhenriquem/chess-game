package game

import (
	"chess-game/internal/model"
	"chess-game/internal/pkg/format"
	"chess-game/internal/protocol"
	"fmt"

	"github.com/gorilla/websocket"
)

func GetOne(id string) *model.Game {
	for _, game := range base {
		if game.ID == id {
			return game
		}
	}
	return nil
}

func Run(id string) {
	g := GetOne(id)
	for _, p := range g.Players {
		msg := fmt.Sprintf("You are playing, you are %s", p.Color)
		protocol.SendMessage(p.Client, "initGame", msg, format.ToFormatGame(g))
	}

	for {
		select {
		case player := <-GetOne(id).Desconnect:
			msg := "Outher player desconnected, you win"
			if player.Color == "black" {
				protocol.SendMessage(GetPlayer(g.ID, "white").Client, "desconnected", msg, format.ToFormatGame(g))

				GetPlayer(g.ID, "white").Client.WriteMessage(websocket.CloseMessage, []byte{})
				GetPlayer(g.ID, "white").Client.Close()
			} else {
				protocol.SendMessage(GetPlayer(g.ID, "black").Client, "desconnected", msg, format.ToFormatGame(g))
				GetPlayer(g.ID, "black").Client.WriteMessage(websocket.CloseMessage, []byte{})
				GetPlayer(g.ID, "black").Client.Close()
			}

			break
		}
	}
}
