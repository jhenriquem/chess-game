package model

import (
	"github.com/corentings/chess/v2"
	"github.com/gorilla/websocket"
)

type Game struct {
	ID string

	Board [8][8]string
	Chess *chess.Game

	Desconnect chan *Player

	MoveResult string
	Moves      [][2]string

	Players [2]*Player

	Timer string

	Turn *Player
}

type Player struct {
	Game          *Game
	Client        *websocket.Conn
	Color         string // B or W
	TimeRemaining string
	Moves         []string
	Score         int
}

func (g *Game) GetPlayer(color string) *Player {
	for _, player := range g.Players {
		if player.Color == color {
			return player
		}
	}
	return nil
}
