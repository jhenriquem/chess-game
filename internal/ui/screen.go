package ui

import "fmt"

func Load(board [8][8]string, isTurn bool) {
	ClearScreen()
	Board(board)

	if isTurn {
		fmt.Print("What's your move ? [initial position] [final position] :")
	}
}
