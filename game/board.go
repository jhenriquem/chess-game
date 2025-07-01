package game

import "fmt"

func (g *Game) InitBoard() {
	var board [8][8]*Pieces

	majorPieces := []string{"Rook", "Knight", "Bishop", "Queen", "King", "Bishop", "Knight", "Rook"}

	columnToLetter := func(col int) string {
		return string(rune('a' + col))
	}

	for col, piece := range majorPieces {
		board[0][col] = &Pieces{
			Location: columnToLetter(col) + "8",
			Color:    g.BlackPlayer.ColorPieces,
			Piece:    piece,
			Player:   g.BlackPlayer,
		}
		board[1][col] = &Pieces{
			Location: columnToLetter(col) + "7",
			Color:    g.BlackPlayer.ColorPieces,
			Piece:    "Pawn",
			Player:   g.BlackPlayer,
		}
	}

	for col, piece := range majorPieces {
		board[7][col] = &Pieces{
			Location: columnToLetter(col) + "1",
			Color:    g.WhitePlayer.ColorPieces,
			Piece:    piece,
			Player:   g.WhitePlayer,
		}
		board[6][col] = &Pieces{
			Location: columnToLetter(col) + "2",
			Color:    g.WhitePlayer.ColorPieces,
			Piece:    "Pawn",
			Player:   g.WhitePlayer,
		}
	}

	for row := 2; row <= 5; row++ {
		for col := 0; col < 8; col++ {
			board[row][col] = &Pieces{
				Location: columnToLetter(col) + fmt.Sprintf("%d", 8-row),
				Color:    "",
				Piece:    "",
			}
		}
	}

	g.Board = board
}
