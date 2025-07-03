package game

import (
	"github.com/gorilla/websocket"
)

func New(p1, p2 *websocket.Conn) *Game {
	g := &Game{
		Players: [2]*Player{newPlayer(p1, "white"), newPlayer(p2, "black")},
		Timer:   "10min",

		MovePlayer: make(chan *Player),
		Desconnect: make(chan *Player),
		Moves:      [][2]string{},
	}

	// Sets the white player as the first to play
	g.Turn = g.GetWhite()

	// Sets games pointers for players
	g.GetWhite().Game = g
	g.GetBlack().Game = g

	// Generates the board linked to the players
	g.InitBoard()

	return g
}
