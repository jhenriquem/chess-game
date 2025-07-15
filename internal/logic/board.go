package logic

import (
	"chess-game/internal/model"
	"chess-game/internal/pkg/pieces"
	"chess-game/internal/ui"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func returnPiece(piece rune) string {
	color := "W"
	if unicode.IsLower(piece) {
		color = "B"
	}

	pieceName := ""
	switch strings.ToLower(string(piece)) {
	case "r":
		pieceName = "rook"
	case "n":
		pieceName = "knight"
	case "b":
		pieceName = "bishop"
	case "k":
		pieceName = "king"
	case "q":
		pieceName = "queen"
	case "p":
		pieceName = "pawn"
	default:
		return "  "
	}
	return pieces.Icons[color+pieceName]
}

func UpdateBoard(game *model.Game) {
	fmt.Println(game.Chess.Position().Board())

	FormatFEN(game)
	ui.Board(game.Board)
}

func FormatFEN(game *model.Game) {
	// rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR
	FEN := game.Chess.Position().Board().String()

	sliceFEN := strings.Split(FEN, "/")
	for l, line := range sliceFEN {
		colomnAmount := 0

		for _, piece := range line {
			if value, err := strconv.Atoi(string(piece)); err == nil {
				for i := 0; i < value; i++ {
					game.Board[l][colomnAmount] = " "
					colomnAmount++
				}
			} else {
				game.Board[l][colomnAmount] = returnPiece(piece)
				colomnAmount++
			}
		}
	}
}
