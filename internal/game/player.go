package game

import (
	"chess-game/model"
	"time"

	"github.com/gorilla/websocket"
)

func FindPlayerByConn(conn *websocket.Conn, game *model.Game) *model.Player {
	for _, player := range game.Players {
		if player.Client == conn {
			return player
		}
	}
	return nil
}

func newPlayer(client *websocket.Conn, color string) *model.Player {
	return &model.Player{
		Client:   client,
		Color:    color,
		Moves:    []string{},
		Score:    0,
		TimeLeft: 15 * time.Minute,
	}
}
