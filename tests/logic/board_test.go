package logic

import (
	"chess-game/internal/logic"
	"testing"

	"github.com/corentings/chess/v2"
)

func TestReturnPiece(t *testing.T) {
	tests := []struct {
		input    rune
		expected string
	}{
		{'k', " ♔ "},
		{'q', " ♕ "},
		{'r', " ♖ "},
		{'b', " ♗ "},
		{'n', " ♘ "},
		{'p', " ♙ "},

		{'K', " 󰡗 "},
		{'Q', " 󰡚 "},
		{'R', " 󰡛 "},
		{'B', " 󰡜 "},
		{'N', " 󰡘 "},
		{'P', " 󰡙 "},
		{'x', " "}, // símbolo inválido
	}

	for _, tt := range tests {
		result := logic.ReturnPiece(tt.input)
		if result != tt.expected {
			t.Errorf("returnPiece(%q) = %q; want %q", tt.input, result, tt.expected)
		} else {
			t.Log("PASS (ReturnPiece)")
		}
	}
}

func TestUpdateBoard(t *testing.T) {
	game := chess.NewGame()
	board := logic.UpdateBoard(game.Position().Board().String())

	if board[0][0] != " ♖ " || board[7][4] != " 󰡗 " {
		t.Errorf("UpdateBoard returned incorrect pieces: got [%v][%v]", board[0][0], board[7][4])
	} else {
		t.Log("PASS (UpdateBoard)")
	}

	// Verifica se as casas vazias realmente estão com " "
	for i := 2; i <= 5; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] != " " {
				t.Errorf("Expected empty square at [%d][%d], got %q", i, j, board[i][j])
			}
		}
	}
}
