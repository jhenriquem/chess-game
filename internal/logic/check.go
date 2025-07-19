package logic

import (
	"chess-game/internal/model"
	"strings"
)

func CheckActionOfMoves(game *model.Game) string {
	gameString := game.Chess.String()

	str := strings.ReplaceAll(gameString, " ", "")

	switch str[len(str)-2] {
	case '+':
		return "CHECK"
	case '#':
		return "CHECKMATE"
	default:
		return ""
	}
}
