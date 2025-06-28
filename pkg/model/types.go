package model

import "github.com/gorilla/websocket"

type Player struct {
	Client        *websocket.Conn
	Name          string
	ColorPieces   string
	TimeRemaining string
}

type Pieces struct {
	Piece    string
	Location string
	Color    string
	Player   *Player // Id of the player to whom this piece belongs
}
