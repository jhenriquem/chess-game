package logic

import (
	"chess-game/internal/logic"
	"testing"
)

func TestForVerifyResultsOfMoves(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"----+.", "Is check"},
		{"----#.", "Is checkmate"},
		{"----.", ""},
	}
	for _, tt := range tests {
		result := logic.VerifyResultsOfMoves(tt.input)
		if result != tt.expected {
			t.Errorf("got %q; want %q", result, tt.expected)
		}
	}
}
