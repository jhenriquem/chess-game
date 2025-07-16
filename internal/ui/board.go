package ui

import (
	"fmt"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func inverterBoard(arr *[8][8]string) {
	start := 0
	end := 7

	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

func Board(board [8][8]string, color string) {
	// Cabeçalho superior
	fmt.Println("    ┌─────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┐")

	if color == "black" {
		inverterBoard(&board)
	}

	for i, row := range board {
		if color == "black" {
			fmt.Printf("  %d  │", i+1) // Números das linhas (de 8 a 1)
		} else {
			fmt.Printf("  %d  │", 8-i) // Números das linhas (de 8 a 1)
		}
		for c, piece := range row {
			if piece == " " {
				fmt.Print("     │")
			} else {
				fmt.Printf(" %s │", row[7-c])
			}
		}
		// Linha horizontal de separação entre as linhas do tabuleiro
		if i < 7 {
			fmt.Println()
			fmt.Println("     ├─────┼─────┼─────┼─────┼─────┼─────┼─────┼─────┤")
		} else {
			fmt.Println()
		}
	}

	// Rodapé com letras das colunas
	fmt.Println("     └─────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┘")
	if color == "black" {
		fmt.Println("        h     g     f     e     d     c     b     a")
	} else {
		fmt.Println("        a     b     c     d     e     f     g     h")
	}
}
