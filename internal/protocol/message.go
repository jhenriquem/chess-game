package protocol

import "chess-game/model"

type Message struct {
	TypeInfo string // INIT, DESCONNECTED, TIMEOUT, TURN, WAIT, PLAYER_FOUND
	Info     string `json:"info"`
	IsTurn   bool
	Game     model.Protogame `json:"game,omitempty"`
}
