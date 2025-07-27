package logic

import (
	"chess-game/model"
	"fmt"
	"strings"

	"github.com/corentings/chess/v2"
)

func MovesLogic(msg model.ClientMessage, game *model.Game) (string, error) {
	if msg.Type == "move" {
		move := strings.Join(msg.Move, "")

		mv, err := ValidMove(move, game)
		if err != nil {
			return "", err
		}

		AddMove(mv, game)

		game.Board = UpdateBoard(game.Chess.Position().Board().String())

		ChangeTurn(game)

		result := VerifyResultsOfMoves(game.Chess.String())

		return result, nil
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

func AddMove(move string, game *model.Game) {
	if len(game.Moves) == 0 {
		game.Moves = append(game.Moves, [2]string{})
	}

	if game.Moves[len(game.Moves)-1][1] != "" {
		game.Moves = append(game.Moves, [2]string{move, ""})
	} else {
		game.Moves[len(game.Moves)-1][1] = move
	}
}

func ValidMove(move string, game *model.Game) (string, error) {
	validMoves := game.Chess.Position().ValidMoves()

	for _, m := range validMoves {
		if m.String() == move {
			if err := game.Chess.PushNotationMove(m.String(), chess.UCINotation{}, nil); err != nil {
				return "", err
			}
			return m.String(), nil
		}
	}

	return "", fmt.Errorf("Invalid move")
}
