package game

import (
	"chess-game/internal/model"

	"github.com/gorilla/websocket"
)

func GetPlayer(id, color string) *model.Player {
	for _, player := range GetOne(id).Players {
		if player.Color == color {
			return player
		}
	}
	return nil
}

func newPlayer(client *websocket.Conn, color string) *model.Player {
	return &model.Player{
		Client: client,
		Color:  color,
		Moves:  []string{},
		Score:  0,
	}
}
