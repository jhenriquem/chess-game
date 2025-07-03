package game

import (
	"chess-game/pkg/protocol"

	"github.com/gorilla/websocket"
)

func (p *Player) ToFormatForJson() protocol.Player {
	return protocol.Player{
		Color:         p.Color,
		Moves:         p.Moves,
		TimeRemaining: p.TimeRemaining,
		Score:         p.Score,
	}
}

func newPlayer(client *websocket.Conn, color string) *Player {
	return &Player{
		Client: client,
		Color:  color,
		Moves:  []string{},
		Score:  0,
	}
}
