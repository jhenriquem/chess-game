package game

import (
	"github.com/corentings/chess/v2"
)

func ColorName(c chess.Color) string {
	if c == chess.White {
		return "White"
	}
	return "Black"
}
