package ui

import (
	"chess-game/internal/protocol"
	"chess-game/model"
	"fmt"
)

func Load(data protocol.Message, player *model.PlayerFormat) {
	ClearScreen()

	// Adicionar o tempo e os nomes dos jogadores
	// Adicionar campos para o ultimo movimento e indicação de turno
	// Melhorar a logica de input de movimentos( bobbletea, net )

	if data.Game.Turn != "" {
		Board(data.Game.Board, player.Color)
	}

	if data.IsTurn {
		fmt.Print("\n 🔵 What's your move ? [initial][final] :")
	}
}
