package format

import (
	"chess-game/model"
)

func ToFormatGame(g *model.Game, player *model.Player) model.Protogame {
	return model.Protogame{
		Board:         g.Board,
		Moves:         g.Moves,
		Player:        ToFormatPlayer(player),
		Timer:         g.Timer,
		Turn:          g.Turn.Color,
		MoveResult:    g.MoveResult,
		Outcome:       g.Chess.Outcome().String(),
		GameString:    g.Chess.String(),
		OutcomeMethod: g.Chess.Method().String(),
	}
}
