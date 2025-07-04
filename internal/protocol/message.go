package protocol

import "chess-game/internal/model"

type Message struct {
	Info string          `json:"info"`
	Game model.Protogame `json:"game,omitempty"`
}
