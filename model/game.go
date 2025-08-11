package model

import (
	"github.com/corentings/chess/v2"
)

type Game struct {
	Chess chess.Game

	Players []*Player
}
