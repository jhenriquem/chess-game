package utils

import "fmt"

func Introdution() {
	fmt.Println("♟️ Welcome to the Game")
	fmt.Println("")
	fmt.Println("📌 This is a terminal chess game.")
	fmt.Println("🔹 When you type [Y] the game will start.")
	fmt.Println("🔹 We will find another player and start the game.")
	fmt.Println("🔹 Warning messages will appear above the board, in green🟢.")
	fmt.Println("")
	fmt.Println("💡 To send your move, just type it and press enter. The moves follow the following formatting")
	fmt.Println("🔹 Piece - [x] - column (a,b,c,..,h) - row (1,2,...8)")
	fmt.Println("🔹 x -> indicates whether your move will perform a capture")
	fmt.Println("")
	fmt.Println("❗ If you want to castl, just type:")
	fmt.Println("🔹 O-O or O-O-O")
}
