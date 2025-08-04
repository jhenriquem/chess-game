package game

import (
	"strconv"
	"strings"
)

func NewEmptyBoard() [8][8]string {
	var board [8][8]string = [8][8]string{{}}

	piecesMajor := []string{"r", "n", "b", "q", "k", "b", "n", "r"}
	for col, piece := range piecesMajor {
		board[0][col] = piece
		board[1][col] = "p"

		board[6][col] = "P"
		board[7][col] = strings.ToUpper(piece)
	}

	for row := 2; row <= 5; row++ {
		for col := 0; col < 8; col++ {
			board[row][col] = " "
		}
	}

	return board
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
				board[l][colomn] = string(piece)
				colomn++
			}
		}
	}

	return board
}
