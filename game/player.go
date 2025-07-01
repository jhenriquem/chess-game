package game

import (
	"github.com/gorilla/websocket"
)

func newPlayer(client *websocket.Conn, color string) *Player {
	return &Player{
		Client:      client,
		ColorPieces: color,
	}
}
