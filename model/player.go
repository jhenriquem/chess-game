package model

import (
	"encoding/json"
	"net"
	"time"

	"github.com/corentings/chess/v2"
)

type Player struct {
	Conn     net.Conn
	Encoder  *json.Encoder
	Decoder  *json.Decoder
	Name     string
	Color    chess.Color
	Timeleft time.Duration
}
