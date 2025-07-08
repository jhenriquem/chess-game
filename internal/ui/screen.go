package ui

import "fmt"

func Load(board [8][8]string) {
	ClearScreen()
	Board(board)
	fmt.Print("What's your move ? (e4)")
}
