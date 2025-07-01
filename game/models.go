package game

import (
	"github.com/gorilla/websocket"
)

type Game struct {
	Board [8][8]*Pieces

	Desconnect chan *Player

	MovePlayer chan *Player
	Moves      []string

	Players [2]*Player

	Timer string

	Turn *Player
}

type Player struct {
	Game          *Game
	Client        *websocket.Conn
	Color         string
	TimeRemaining string
}

type Pieces struct {
	Piece    string
	Location string
	Color    string
	Player   *Player
}
