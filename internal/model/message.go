package model

type ClientMessage struct {
	Type string // move, draw, giveUp
	Move []string
}
