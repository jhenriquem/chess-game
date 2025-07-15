package format

import (
	"chess-game/internal/model"
)

func ToFormatGame(g *model.Game) model.Protogame {
	return model.Protogame{
		Board:   g.Board,
		Moves:   g.Moves,
		Players: [2]model.Protoplayer{ToFormatPlayer(g.Players[0]), ToFormatPlayer(g.Players[1])},
		Timer:   g.Timer,
		Turn:    g.Turn.Color,
	}
}
