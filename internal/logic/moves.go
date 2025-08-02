package logic

import (
	"chess-game/model"
	"fmt"

	"github.com/corentings/chess/v2"
)

func Moves(msg model.ClientMessage, game *model.Game) error {
	move := msg.Move

	mv, err := ValidMove(move, game)
	if err != nil {
		return err
	}

	AddMove(mv, game)

	VerifyResultsOfMoves(game)

	return nil
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
		fmt.Print(m.String() + "´ ")
		if m.String() == move {
			if err := game.Chess.PushNotationMove(m.String(), chess.UCINotation{}, nil); err != nil {
				return "", err
			}
			return m.String(), nil
		}
	}

	return "", fmt.Errorf("Invalid move")
}
