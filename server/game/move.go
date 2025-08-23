package game

import (
	"chess-game/model"
	"fmt"
	"strings"

	"github.com/corentings/chess/v2"
)

func ReturnLastMove(game *model.Game) string {
	moves := strings.Split(game.Chess.String(), ".")
	last := fmt.Sprintf("%d. %s", len(moves), moves[len(moves)-1])

	return last
}

func ValidMove(move string, c *chess.Game) (bool, error) {
	for _, validMove := range c.Position().ValidMoves() {
		if validMove.String() == move {
			err := c.PushNotationMove(move, chess.UCINotation{}, nil)
			if err != nil {
				return false, err
			}
			return true, nil
		}
	}
	return false, fmt.Errorf("Invalid Move")
}
