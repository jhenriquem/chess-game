package game

import (
	"chess-game/model"
	"fmt"
	"time"

	"github.com/corentings/chess/v2"
)

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
				White: model.PlayerFormat{
					Color:    game.Players[0].Color,
					Name:     game.Players[0].Name,
					Timeleft: game.Players[0].Timeleft,
				},
				Black: model.PlayerFormat{
					Color:    game.Players[1].Color,
					Timeleft: game.Players[1].Timeleft,
					Name:     game.Players[1].Name,
				},
				FEN:      game.Chess.Position().Board().String(),
				Message:  fmt.Sprintf("You are %s", ColorName(p.Color)),
				LastMove: ReturnLastMove(&game),
			},
		}
		p.Encoder.Encode(msg)
	}

	Run(&game)
}
