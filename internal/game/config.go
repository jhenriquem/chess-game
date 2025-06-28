package game

import (
	"chess-game_cli/pkg/model"
)

type GameStruct struct {
	Board [8][8]model.Pieces
	Moves []string
}

func New() *GameStruct {
	return &GameStruct{
		Board: [8][8]model.Pieces{},
	}
}

// func (g *GameStruct) GetPlayer
