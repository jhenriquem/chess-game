package protocol

import "chess-game/model"

type Message struct {
	TypeInfo string // INIT, DESCONNECTED, WAIT_TIMEOUT, TIMEOUT, TURN, WAIT, PLAYER_FOUND
	Info     string `json:"info"`
	IsTurn   bool
	Game     model.GameFormat `json:"game,omitempty"`
}
