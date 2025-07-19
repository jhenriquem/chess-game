package game

import (
	"chess-game/internal/model"

	"github.com/corentings/chess/v2"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var base []*model.Game = []*model.Game{}

func New(p1, p2 *websocket.Conn) *model.Game {
	var game model.Game = model.Game{
		ID: uuid.NewString(),

		Board:   [8][8]string{},
		Players: [2]*model.Player{newPlayer(p1, "white"), newPlayer(p2, "black")},
		Timer:   "10min",

		Chess: chess.NewGame(),

		Desconnect: make(chan *model.Player),
		Moves:      [][2]string{{}},
	}

	base = append(base, &game)

	// Sets the white player as the first to play
	game.Turn = GetPlayer(&game, "white")

	// Sets games pointers for players
	GetPlayer(&game, "white").Game = &game
	GetPlayer(&game, "black").Game = &game

	// Generates the board linked to the players
	SetupBoard(&game)

	return &game
}
