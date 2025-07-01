package game

import (
	"github.com/gorilla/websocket"
)

func (p *Player) SendInfo(msg string) {
	p.Client.WriteJSON(map[string]string{"info": msg})
}

func newPlayer(client *websocket.Conn, color string) *Player {
	return &Player{
		Client: client,
		Color:  color,
	}
}
