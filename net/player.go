package net

import (
	"chess-game/model"
	"log"
	"time"
)

func ReadPlayerMessage(player *model.Player, moveChan chan model.Message) bool {
	var playerMessage model.Message

	player.Conn.SetReadDeadline(time.Now().Add(300 * time.Second))

	if err := player.Conn.ReadJSON(&playerMessage); err != nil {
		log.Printf("\nError reading player(%s) message (%s)", player.Color, err.Error())

		close(moveChan)

		return false
	}

	moveChan <- playerMessage
	return true
}
