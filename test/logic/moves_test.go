package logic

import (
	"chess-game/internal/logic"
	"chess-game/model"
	"errors"
	"testing"

	"github.com/corentings/chess/v2"
)

func TestValidMoves(t *testing.T) {
	testsMatrix := []struct {
		input    model.ClientMessage
		expected error
	}{
		{input: model.ClientMessage{Move: "e2e4"}, expected: nil},
		{input: model.ClientMessage{Move: "e7e5"}, expected: nil},
		{input: model.ClientMessage{Move: "d2d4"}, expected: nil},
		{input: model.ClientMessage{Move: "e5d4"}, expected: nil},
		{input: model.ClientMessage{Move: "h2h5"}, expected: errors.New("Invalid move")},
		{input: model.ClientMessage{Move: "g1h5"}, expected: errors.New("Invalid move")},
	}

	gameMock := model.Game{
		Chess: chess.NewGame(),
		Moves: [][2]string{{}},
	}

	for _, tt := range testsMatrix {
		err := logic.Moves(tt.input, &gameMock)
		if err != tt.expected && err.Error() != tt.expected.Error() {
			t.Errorf("Move %s , expected %v, got %v", tt.input.Move, tt.expected, err)
		}
	}
}
