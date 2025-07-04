package model

import (
	"github.com/gorilla/websocket"
)

type Game struct {
	ID string

	Board [8][8]*Pieces

	Desconnect chan *Player

	MovePlayer chan *Player
	Moves      [][2]string

	Players [2]*Player

	Timer string

	Turn *Player
}

type Player struct {
	Game          *Game
	Client        *websocket.Conn
	Color         string
	TimeRemaining string
	Moves         []string
	Score         int
}

type Pieces struct {
	Piece    string
	Location string
	Color    string
	Player   *Player
	Icon     string
}
