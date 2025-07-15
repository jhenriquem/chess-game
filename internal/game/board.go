package game

import (
	"chess-game/internal/model"
	"chess-game/internal/pkg/pieces"
)

func SetupBoard(game *model.Game) {
	var board [8][8]string = [8][8]string{{}}

	for col, piece := range pieces.Major {
		board[0][col] = pieces.Icons["B"+piece]
		board[1][col] = pieces.Icons["Bpawn"]

		board[6][col] = pieces.Icons["Wpawn"]
		board[7][col] = pieces.Icons["W"+piece]
	}

	for row := 2; row <= 5; row++ {
		for col := 0; col < 8; col++ {
			board[row][col] = " "
		}
	}

	game.Board = board
}
