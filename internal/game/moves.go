package game

import (
	"fmt"

	"github.com/corentings/chess/v2"
)

func ValidMove(move string, game *chess.Game) (string, bool) {
	validMoves := game.ValidMoves()

	for _, m := range validMoves {
		if m.String() == move {
			fmt.Println(m.String(), move)

			game.Move(&m, nil)
			return m.String(), true
		}
	}
	return move, false
}
