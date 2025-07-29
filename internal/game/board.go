package game

import (
	"chess-game/pkg/pieces"
	"strconv"
	"strings"
	"unicode"
)

func NewEmptyBoard() [8][8]string {
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

	return board
}

func returnPiece(piece rune) string {
	color := "W"

	if unicode.IsLower(piece) {
		color = "B"
	}

	pieceName := ""

	switch strings.ToLower(string(piece)) {
	case "r":
		pieceName = "rook"
	case "n":
		pieceName = "knight"
	case "b":
		pieceName = "bishop"
	case "k":
		pieceName = "king"
	case "q":
		pieceName = "queen"
	case "p":
		pieceName = "pawn"
	default:
		return " "
	}

	return pieces.Icons[color+pieceName]
}

func UpdateBoard(FEN string) [8][8]string {
	// rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR

	board := [8][8]string{}
	sliceFEN := strings.Split(FEN, "/")
	for l, line := range sliceFEN {
		colomn := 0

		for _, piece := range line {
			if value, err := strconv.Atoi(string(piece)); err == nil {
				for i := 0; i < value; i++ {
					board[l][colomn] = " "
					colomn++
				}
			} else {
				board[l][colomn] = returnPiece(piece)
				colomn++
			}
		}
	}

	return board
}
