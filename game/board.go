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
			Color:    g.GetBlack().Color,
			Piece:    piece,
			Player:   g.GetBlack(),
		}
		board[1][col] = &Pieces{
			Location: columnToLetter(col) + "7",
			Color:    g.GetBlack().Color,
			Piece:    "Pawn",
			Player:   g.GetBlack(),
		}
	}

	for col, piece := range majorPieces {
		board[7][col] = &Pieces{
			Location: columnToLetter(col) + "1",
			Color:    g.GetWhite().Color,
			Piece:    piece,
			Player:   g.GetWhite(),
		}
		board[6][col] = &Pieces{
			Location: columnToLetter(col) + "2",
			Color:    g.GetWhite().Color,
			Piece:    "Pawn",
			Player:   g.GetWhite(),
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
