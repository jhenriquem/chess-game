package ui

import "fmt"

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func Board(board [8][8]string) {
	// Cabeçalho superior
	fmt.Println("    ┌─────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┐")

	for i, row := range board {
		fmt.Printf(" %d  │", 8-i) // Números das linhas (de 8 a 1)
		for _, piece := range row {
			if piece == "" {
				fmt.Print("     │")
			} else {
				fmt.Printf(" %s │", piece)
			}
		}
		// Linha horizontal de separação entre as linhas do tabuleiro
		if i < 7 {
			fmt.Println()
			fmt.Println("    ├─────┼─────┼─────┼─────┼─────┼─────┼─────┼─────┤")
		} else {
			fmt.Println()
		}
	}

	// Rodapé com letras das colunas
	fmt.Println("    └─────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┘")
	fmt.Println("       a     b     c     d     e     f     g     h")
}
