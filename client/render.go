package client

import "fmt"

func RenderBoard(board [8][8]string) {
	for i, row := range board {
		fmt.Printf("%d ", 8-i) // Mostrar linhas invertidas
		for _, piece := range row {
			if piece == "" {
				fmt.Print("|   ")
			} else {
				fmt.Printf("|%s", piece)
			}
		}
		fmt.Println("|")
	}
	fmt.Println("    a   b   c   d   e   f   g   h")
	fmt.Println()
}
