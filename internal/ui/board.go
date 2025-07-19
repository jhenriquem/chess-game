package ui

import (
	"fmt"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func inverterBoard(arr [8][8]string) [8][8]string {
	var board [8][8]string
	start := 0
	end := 7

	for start < end {
		board[start], board[end] = arr[end], arr[start]
		start++
		end--
	}

	for i, line := range board {

		start = 0
		end = 7

		for range line {
			if start < end {
				board[i][start], board[i][end] = board[i][end], board[i][start]
				start++
				end--
			}
		}
	}

	return board
}

func Board(gameBoard [8][8]string, color string) {
	board := gameBoard

	if color == "black" {
		board = inverterBoard(board)
	}

	fmt.Println("    ┌─────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┐")

	for i, row := range board {
		lineNumber := 8 - i
		if color == "black" {
			lineNumber = i + 1
		}
		fmt.Printf("  %d  │", lineNumber)

		for _, piece := range row {
			if piece == " " {
				fmt.Print("     │")
			} else {
				fmt.Printf(" %s │", piece)
			}
		}

		if i < 7 {
			fmt.Println()
			fmt.Println("     ├─────┼─────┼─────┼─────┼─────┼─────┼─────┼─────┤")
		} else {
			fmt.Println()
		}
	}

	fmt.Println("     └─────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┘")
	if color == "black" {
		fmt.Println("        h     g     f     e     d     c     b     a")
	} else {
		fmt.Println("        a     b     c     d     e     f     g     h")
	}
}
