package logic

import (
	"chess-game/internal/logic"
	"chess-game/model"
	"testing"

	"github.com/corentings/chess/v2"
)

func TestOfValidMoves(t *testing.T) {
	game := model.Game{
		Chess: chess.NewGame(),
	}

	tests := []struct {
		move     string
		expected error
	}{
		{"e2e4", nil},
		{"g8f6", nil},
		{"e3e4", nil},
	}

	for i, tt := range tests {
		result, err := logic.ValidMove(tt.move, &game)
		if err != tt.expected {
			if i != 2 {
				t.Errorf("validMove(%q) = %q; want %q", tt.move, result, tt.expected)
			}
		}
	}
}
