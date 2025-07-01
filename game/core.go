package game

import (
	"github.com/gorilla/websocket"
)

func New(p1, p2 *websocket.Conn) *Game {
	g := &Game{
		BlackPlayer: newPlayer(p2, "black"),
		WhitePlayer: newPlayer(p1, "white"),
		Timer:       "10min",
		MovePlayer:  make(chan *Player),
		Desconnect:  make(chan *Player),
		Moves:       []string{},
	}

	// Sets the white player as the first to play
	g.Turn = g.WhitePlayer

	// Sets games pointers for players
	g.BlackPlayer.Game = g
	g.WhitePlayer.Game = g

	// Generates the board linked to the players
	g.InitBoard()

	return g
}
