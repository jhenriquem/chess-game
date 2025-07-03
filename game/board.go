package game

import (
	"chess-game/pkg/pieces"
	"fmt"
)

func (g *Game) InitBoard() {
	var board [8][8]*Pieces

	columnToLetter := func(col int) string {
		return string(rune('a' + col))
	}

	for col, piece := range pieces.Major {
		board[0][col] = &Pieces{
			Location: columnToLetter(col) + "8",
			Color:    g.GetBlack().Color,
			Piece:    piece,
			Icon:     pieces.Icons["B"+piece],
			Player:   g.GetBlack(),
		}
		board[1][col] = &Pieces{
			Location: columnToLetter(col) + "7",
			Color:    g.GetBlack().Color,
			Piece:    "pawn",
			Icon:     pieces.Icons["Bpawn"],
			Player:   g.GetBlack(),
		}
	}

	for col, piece := range pieces.Major {
		board[7][col] = &Pieces{
			Location: columnToLetter(col) + "1",
			Color:    g.GetWhite().Color,
			Piece:    piece,
			Icon:     pieces.Icons["W"+piece],
			Player:   g.GetWhite(),
		}
		board[6][col] = &Pieces{
			Location: columnToLetter(col) + "2",
			Color:    g.GetWhite().Color,
			Piece:    "Pawn",
			Icon:     pieces.Icons["Wpawn"],
			Player:   g.GetWhite(),
		}
	}

	for row := 2; row <= 5; row++ {
		for col := 0; col < 8; col++ {
			board[row][col] = &Pieces{
				Location: columnToLetter(col) + fmt.Sprintf("%d", 8-row),
				Color:    "",
				Piece:    "",
			}
		}
	}

	g.Board = board
}

func (g *Game) ToFormatBoard() [8][8]string {
	board := [8][8]string{}
	for i, row := range g.Board {
		for n, column := range row {
			board[i][n] = column.Icon
		}
	}
	return board
}
