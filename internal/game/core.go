package game

import (
	"chess-game/model"

	"github.com/corentings/chess/v2"
	"github.com/gorilla/websocket"
)

func New(p1, p2 *websocket.Conn) *model.Game {
	var game model.Game = model.Game{
		Board: NewEmptyBoard(),

		Players: [2]*model.Player{newPlayer(p1, "W"), newPlayer(p2, "B")},

		Timer: "10min",

		Chess: chess.NewGame(),

		Desconnect: make(chan *model.Player),
		Moves:      [][2]string{{}},

		MoveChan: make(chan string),
	}

	// Sets the white player as the first to play
	game.CurrentPlayer = game.GetPlayer("W")

	// Sets games pointers for players
	game.GetPlayer("W").Game = &game
	game.GetPlayer("B").Game = &game

	return &game
}
