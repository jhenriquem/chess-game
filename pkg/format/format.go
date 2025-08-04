package format

import (
	"chess-game/model"
)

func FormatPlayer(p *model.Player) model.PlayerFormat {
	return model.PlayerFormat{
		Color:    p.Color,
		Moves:    p.Moves,
		TimeLeft: p.TimeLeft,
		Score:    p.Score,
		Name:     p.Name,
	}
}

func Game(game *model.Game) model.GameFormat {
	return model.GameFormat{
		Board:         game.Board,
		Moves:         game.Moves,
		Players:       []model.PlayerFormat{FormatPlayer(game.Players[0]), FormatPlayer(game.Players[1])},
		Timer:         game.Timer,
		Turn:          game.CurrentPlayer.Color,
		IsCheck:       game.IsCheck,
		IsMate:        game.IsMate,
		Outcome:       game.Chess.Outcome().String(),
		GameString:    game.Chess.String(),
		OutcomeMethod: game.Chess.Method().String(),
	}
}
