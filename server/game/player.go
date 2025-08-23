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

func ReadPlayerMessage(player *model.Player, moveChan chan model.Message) bool {
	var playerMessage model.Message

	player.Conn.SetReadDeadline(time.Now().Add(300 * time.Second))

	if err := player.Decoder.Decode(&playerMessage); err != nil {
		log.Printf("\nError reading player(%s) message (%s)", player.Color, err.Error())

		close(moveChan)

		return false
	}
	moveChan <- playerMessage
	return true
}
