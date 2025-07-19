package ui

import (
	"chess-game/internal/protocol"
	"fmt"
)

func Load(data protocol.Message) {
	ClearScreen()

	if data.Info != "" {
		fmt.Printf("\n %s \n ", data.Info)
	}

	if data.Game.Turn != "" {
		Board(data.Game.Board, data.Game.Player.Color)
	}

	ShowMessage(data.Game.MoveResult)

	if data.IsTurn {
		fmt.Print("\n 🔵 What's your move ? [initial][final] :")
	}
}
