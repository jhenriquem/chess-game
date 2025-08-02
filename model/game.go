package model

import (
	"time"

	"github.com/corentings/chess/v2"
	"github.com/gorilla/websocket"
)

type Game struct {
	Board [8][8]string
	Chess *chess.Game

	Desconnect chan *Player
	Timeout    chan *Player

	IsCheck bool
	IsMate  bool

	MoveResult string
	Moves      [][2]string

	Players [2]*Player

	Timer string

	CurrentPlayer *Player
}

type Player struct {
	Game   *Game
	Client *websocket.Conn
	Color  string // B or W

	TimeLeft time.Duration
	Timer    *time.Timer
	StopChan chan struct{}

	Moves []string
	Score int
}

// Return the player other than the current one
func (g *Game) GetAnoutherPlayer() *Player {
	for _, player := range g.Players {
		if player.Color != g.CurrentPlayer.Color {
			return player
		}
	}
	return nil
}

func (g *Game) GetPlayer(color string) *Player {
	for _, player := range g.Players {
		if player.Color == color {
			return player
		}
	}
	return nil
}
