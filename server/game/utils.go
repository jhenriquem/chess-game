package game

import (
	"chess-game/model"
	"strings"

	"github.com/corentings/chess/v2"
)

func ColorName(c chess.Color) string {
	if c == chess.White {
		return "White"
	}
	return "Black"
}

func ReturnLastMove(game *model.Game) string {
	moves := strings.Split(game.Chess.String(), ".")
	last := moves[len(moves)-1]

	return last
}
