package model

type ClientMessage struct {
	Type   string // MOVE, DRAW, RESING
	Move   string
	Player *Player
}
