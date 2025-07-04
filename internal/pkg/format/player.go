package format

import (
	"chess-game/internal/model"
)

func ToFormatPlayer(p *model.Player) model.Protoplayer {
	return model.Protoplayer{
		Color:         p.Color,
		Moves:         p.Moves,
		TimeRemaining: p.TimeRemaining,
		Score:         p.Score,
	}
}
