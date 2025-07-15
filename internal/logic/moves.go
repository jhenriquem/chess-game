package logic

import (
	"chess-game/internal/model"
	"fmt"
	"strings"

	"github.com/corentings/chess/v2"
)

func addMove(move string, game *model.Game) {
	if len(game.Moves) == 0 {
		game.Moves = append(game.Moves, [2]string{})
	}

	if game.Moves[len(game.Moves)-1][1] != "" {
		game.Moves = append(game.Moves, [2]string{move, ""})
	} else {
		game.Moves[len(game.Moves)-1][1] = move
	}
}

func MovesLogic(msg model.ClientMessage, game *model.Game) error {
	if msg.Type == "move" {
		move := strings.Join(msg.Move, "")

		if move, isValid := validMove(move, game); isValid {
			addMove(move, game)
			UpdateBoard(game)
			return nil
		}
	}
	return fmt.Errorf("Invalid move")
}

func validMove(move string, game *model.Game) (string, bool) {
	validMoves := game.Chess.ValidMoves()

	for _, m := range validMoves {
		fmt.Print(m.String())
		if m.String() == move {

			game.Chess.PushNotationMove(m.String(), chess.AlgebraicNotation{}, nil)
			return m.String(), true
		}
	}
	return move, false
}
