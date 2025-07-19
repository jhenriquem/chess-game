package logic

import (
	"chess-game/internal/model"
	"fmt"
	"strings"

	"github.com/corentings/chess/v2"
)

func MovesLogic(msg model.ClientMessage, game *model.Game) (string, error) {
	if msg.Type == "move" {
		move := strings.Join(msg.Move, "")

		if mv, isValid := validMove(move, game); isValid {
			addMove(mv, game)
			UpdateBoard(game)

			ChangeTurn(game)

			result := CheckActionOfMoves(game)

			return result, nil
		} else {
			return "", fmt.Errorf("Invalid move or error: %s", mv)
		}
	}
	return "", nil
}

func ChangeTurn(game *model.Game) {
	if game.Chess.Position().Turn() == 1 {
		game.Turn = game.Players[0]
	} else if game.Chess.Position().Turn() == 2 {
		game.Turn = game.Players[1]
	}
}

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

func validMove(move string, game *model.Game) (string, bool) {
	validMoves := game.Chess.Position().ValidMoves()

	for _, m := range validMoves {
		if m.String() == move {
			if err := game.Chess.PushNotationMove(m.String(), chess.UCINotation{}, nil); err != nil {
				return err.Error(), false
			}
			return m.String(), true
		}
	}
	return move, false
}
