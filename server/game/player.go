package game

import (
	"chess-game/model"
	"encoding/json"
	"log"
	"net"
	"time"
)

func NewPlayer(conn net.Conn) *model.Player {
	enc := json.NewEncoder(conn)
	dec := json.NewDecoder(conn)

	var connMessage model.Message

	conn.SetReadDeadline(time.Now().Add(30 * time.Second))

	if err := dec.Decode(&connMessage); err != nil {
		log.Println("failed to read connection:", err)
		conn.Close()
		return nil
	}

	if connMessage.Type != "CONNECTED" {
		enc.Encode(model.Message{Type: "ERROR", Data: model.Data{FEN: "", Message: "First message must be CONNECTED"}})
		conn.Close()
		return nil
	}

	player := model.Player{
		Conn:    conn,
		Decoder: dec,
		Encoder: enc,
		Name:    connMessage.Data.Message,
	}

	return &player
}
