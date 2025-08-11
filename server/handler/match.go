package handler

import (
	"chess-game/model"
	"chess-game/server/game"
	"log"
)

func Match(p1, p2 *model.Player) {
	log.Printf("New Game: %s (W) vs %s (B)\n", p1.Name, p2.Name)

	// Notifying players that they have encountered a player
	data := model.Data{
		FEN:     "",
		Message: "Player found, starting game",
	}

	msg := model.Message{
		Type: "PLAYER_FOUND",
		Data: data,
	}

	p1.Encoder.Encode(msg)
	p2.Encoder.Encode(msg)

	game.Start(p1, p2)
}
