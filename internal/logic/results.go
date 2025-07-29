package logic

import (
	"chess-game/model"
	"strings"
)

// Functions responsible for verifying the results of movements (check, checkmate)

func IsCheck(result rune) bool {
	if result == '+' {
		return true
	}
	return false
}

func IsMate(result rune) bool {
	if result == '#' {
		return true
	}
	return false
}

func VerifyResultsOfMoves(game *model.Game) {
	gameString := game.Chess.String()
	str := strings.ReplaceAll(gameString, " ", "")
	resultRune := rune(str[len(str)-2])

	game.IsCheck = IsCheck(resultRune)
	game.IsMate = IsMate(resultRune)
}
