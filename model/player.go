package model

import (
	"time"

	"github.com/corentings/chess/v2"
	"github.com/gorilla/websocket"
)

type PlayerFormat struct {
	Name     string
	Color    chess.Color
	Timeleft time.Duration
}

type Player struct {
	Conn     *websocket.Conn
	Name     string
	Color    chess.Color
	Timeleft time.Duration
	Send     chan Message
}
