package game

import (
	"chess-game/internal/model"
	"chess-game/internal/pkg/pieces"
	"fmt"
)

func SetupBoard(id string) {
	var board [8][8]*model.Pieces

	columnToLetter := func(col int) string {
		return string(rune('a' + col))
	}

	for col, piece := range pieces.Major {
		board[0][col] = &model.Pieces{
			Location: columnToLetter(col) + "8",
			Color:    "black",
			Piece:    piece,
			Icon:     pieces.Icons["B"+piece],
			Player:   GetPlayer(id, "black"),
		}
		board[1][col] = &model.Pieces{
			Location: columnToLetter(col) + "7",
			Color:    "black",
			Piece:    "pawn",
			Icon:     pieces.Icons["Bpawn"],
			Player:   GetPlayer(id, "black"),
		}
	}

	for col, piece := range pieces.Major {
		board[7][col] = &model.Pieces{
			Location: columnToLetter(col) + "1",
			Color:    "white",
			Piece:    piece,
			Icon:     pieces.Icons["W"+piece],
			Player:   GetPlayer(id, "white"),
		}
		board[6][col] = &model.Pieces{
			Location: columnToLetter(col) + "2",
			Color:    "white",
			Piece:    "Pawn",
			Icon:     pieces.Icons["Wpawn"],
			Player:   GetPlayer(id, "white"),
		}
	}

	for row := 2; row <= 5; row++ {
		for col := 0; col < 8; col++ {
			board[row][col] = &model.Pieces{
				Location: columnToLetter(col) + fmt.Sprintf("%d", 8-row),
				Color:    "",
				Piece:    "",
			}
		}
	}

	GetOne(id).Board = board
}
