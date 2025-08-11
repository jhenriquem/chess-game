package game

import (
	"chess-game/model"
	"fmt"
	"time"

	"github.com/corentings/chess/v2"
)

func ColorName(c chess.Color) string {
	if c == chess.White {
		return "White"
	}
	return "Black"
}

func Start(p1, p2 *model.Player) {
	// Assign colors
	p1.Color = chess.White
	p2.Color = chess.Black

	// Create new game
	game := model.Game{
		Chess:   *chess.NewGame(),
		Players: []*model.Player{p1, p2},
	}

	// Send initial info
	for _, p := range game.Players {
		p.Timeleft = 15 * time.Minute
		msg := model.Message{
			Type: "START",
			Data: model.Data{
				FEN:     "",
				Message: fmt.Sprintf("You are %s", ColorName(p.Color)),
			},
		}
		p.Encoder.Encode(msg)
	}

	Run(&game)
}
