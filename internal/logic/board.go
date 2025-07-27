package logic

import (
	"chess-game/pkg/pieces"
	"strconv"
	"strings"
	"unicode"
)

func ReturnPiece(piece rune) string {
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
				board[l][colomn] = ReturnPiece(piece)
				colomn++
			}
		}
	}

	return board
}
