package game

import (
	"chess-game/internal/model"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var base []*model.Game = []*model.Game{}

func New(p1, p2 *websocket.Conn) string {
	var game model.Game = model.Game{
		ID: uuid.NewString(),

		Players: [2]*model.Player{newPlayer(p1, "white"), newPlayer(p2, "black")},
		Timer:   "10min",

		MovePlayer: make(chan *model.Player),
		Desconnect: make(chan *model.Player),
		Moves:      [][2]string{},
	}

	base = append(base, &game)

	// Sets the white player as the first to play
	game.Turn = GetPlayer(game.ID, "white")

	// Sets games pointers for players
	GetPlayer(game.ID, "white").Game = &game
	GetPlayer(game.ID, "black").Game = &game

	// Generates the board linked to the players
	SetupBoard(game.ID)

	return game.ID
}
