package handler

import (
	"chess-game/model"
)

func HandleClientMessage(game *model.Game, message chan model.ClientMessage, done <-chan struct{}) {
	for {
		select {
		case <-done:
			return

		case msg := <-message:
			if msg.Type == "MOVE" {
				HandlerPlayerMove(game, msg)
			}
		}
	}
}
