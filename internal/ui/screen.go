package ui

import (
	"chess-game/internal/protocol"
	"chess-game/model"
	"fmt"
)

func Load(data protocol.Message, player *model.PlayerFormat) {
	ClearScreen()

	if data.Info != "" {
		fmt.Printf("\n %s \n ", data.Info)
	}

	if data.Game.Turn != "" {
		Board(data.Game.Board, player.Color)
	}

	if data.Game.IsCheck {
		fmt.Println("Is check")
	}

	fmt.Println(player.TimeLeft)

	if data.Game.IsMate {
		fmt.Println("Checkmate")
	}

	if data.IsTurn {
		fmt.Print("\n 🔵 What's your move ? [initial][final] :")
	}
}
