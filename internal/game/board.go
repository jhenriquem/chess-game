package game

import (
	"chess-game_cli/pkg/model"
	"fmt"
)

func (g *GameStruct) InitBoard(p1, p2 *model.Player) {
	var board [8][8]model.Pieces

	majorPieces := []string{"Rook", "Knight", "Bishop", "Queen", "King", "Bishop", "Knight", "Rook"}

	columnToLetter := func(col int) string {
		return string(rune('a' + col))
	}

	for col, piece := range majorPieces {
		board[0][col] = model.Pieces{
			Location: columnToLetter(col) + "8",
			Color:    p2.ColorPieces,
			Piece:    piece,
			Player:   p2,
		}
		board[1][col] = model.Pieces{
			Location: columnToLetter(col) + "7",
			Color:    p2.ColorPieces,
			Piece:    "Pawn",
			Player:   p2,
		}
	}

	for col, piece := range majorPieces {
		board[7][col] = model.Pieces{
			Location: columnToLetter(col) + "1",
			Color:    p1.ColorPieces,
			Piece:    piece,
			Player:   p1,
		}
		board[6][col] = model.Pieces{
			Location: columnToLetter(col) + "2",
			Color:    p1.ColorPieces,
			Piece:    "Pawn",
			Player:   p1,
		}
	}

	for row := 2; row <= 5; row++ {
		for col := 0; col < 8; col++ {
			board[row][col] = model.Pieces{
				Location: columnToLetter(col) + fmt.Sprintf("%d", 8-row),
				Color:    "",
				Piece:    "",
			}
		}
	}

	g.Board = board
}

func (g *GameStruct) GetBoard() [8][8]model.Pieces {
	return g.Board
}

func (g *GameStruct) SetHouse(positon [2]int, piece model.Pieces) {
	g.Board[positon[1]][positon[0]] = piece
}
