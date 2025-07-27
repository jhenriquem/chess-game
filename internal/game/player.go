package game

import (
	"chess-game/model"

	"github.com/gorilla/websocket"
)

func newPlayer(client *websocket.Conn, color string) *model.Player {
	return &model.Player{
		Client: client,
		Color:  color,
		Moves:  []string{},
		Score:  0,
	}
}
