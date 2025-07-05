package protocol

import "chess-game/internal/model"

type Message struct {
	TypeInfo string          // initGame, desconnected, timeout, yourTurn, playerMove, waiting, playerFound
	Info     string          `json:"info"`
	Game     model.Protogame `json:"game,omitempty"`
}
