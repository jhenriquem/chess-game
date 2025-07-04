package format

import (
	"chess-game/internal/model"
)

func ToFormatBoard(game *model.Game) [8][8]string {
	board := [8][8]string{}
	for i, row := range game.Board {
		for n, column := range row {
			board[i][n] = column.Icon
		}
	}
	return board
}
