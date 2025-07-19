package format

import (
	"chess-game/internal/model"
)

func ToFormatGame(g *model.Game, player *model.Player) model.Protogame {
	return model.Protogame{
		Board:      g.Board,
		Moves:      g.Moves,
		Player:     ToFormatPlayer(player),
		Timer:      g.Timer,
		Turn:       g.Turn.Color,
		MoveResult: g.MoveResult,
	}
}
