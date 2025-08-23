package game

import (
	"chess-game/model"
	"fmt"

	"github.com/corentings/chess/v2"
)

func EndGameByResign(player, oponnent *model.Player, game *model.Game) {
	game.Chess.Resign(player.Color)

	oponnent.Encoder.Encode(model.Message{
		Type: "END",
		Data: model.Data{
			FEN:      game.Chess.Position().Board().String(),
			LastMove: ReturnLastMove(game),
			Message:  "Game Over - Your oponnent has given up, you won ",
			Status:   game.Chess.Method().String(),
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
		},
	})
}

func EndGameByTimeUp(player, oponnent *model.Player, game *model.Game) {
	game.Chess.Resign(player.Color)
	player.Encoder.Encode(model.Message{
		Type: "END",
		Data: model.Data{
			Message:  fmt.Sprintf("Game Over - Your time is up (%s)", game.Chess.Outcome()),
			Status:   "Time is up",
			LastMove: ReturnLastMove(game),
			FEN:      game.Chess.Position().Board().String(),
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
		},
	})

	oponnent.Encoder.Encode(model.Message{
		Type: "END",
		Data: model.Data{
			FEN:      game.Chess.Position().Board().String(),
			LastMove: ReturnLastMove(game),
			Message:  fmt.Sprintf("Game Over - Your oponnent's time is up (%s)", game.Chess.Outcome()),
			Status:   "Time is up",
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
		},
	})
}

func EndGame(player, oponnent *model.Player, game *model.Game) {
	outcome := game.Chess.Outcome()
	var msgCurrent, msgOpponent string

	switch outcome {
	case chess.WhiteWon:
		if player.Color == chess.White {
			msgCurrent, msgOpponent = "Game Over - You won", "Game Over - You lost"
		} else {
			msgCurrent, msgOpponent = "Game Over - You lost", "Game Over - You won"
		}
	case chess.BlackWon:
		if player.Color == chess.Black {
			msgCurrent, msgOpponent = "Game Over - You won", "Game Over - You lost"
		} else {
			msgCurrent, msgOpponent = "Game Over - You lost", "Game Over - You won"
		}
	case chess.Draw:
		msgCurrent, msgOpponent = "Draw", "Draw"
	}

	msgCurrent = fmt.Sprintf("%s by %s - %s", msgCurrent, game.Chess.Method(), outcome.String())
	msgOpponent = fmt.Sprintf("%s by %s - %s", msgOpponent, game.Chess.Method(), outcome.String())

	player.Encoder.Encode(model.Message{
		Type: "END",
		Data: model.Data{
			Message:  msgCurrent,
			Status:   game.Chess.Method().String(),
			LastMove: ReturnLastMove(game),
			FEN:      game.Chess.Position().Board().String(),
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
		},
	})

	oponnent.Encoder.Encode(model.Message{
		Type: "END",
		Data: model.Data{
			FEN:      game.Chess.Position().Board().String(),
			LastMove: ReturnLastMove(game),
			Message:  msgOpponent,
			Status:   game.Chess.Method().String(),
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
		},
	})
}
