package handler

import (
	g "chess-game/internal/game"
	"chess-game/internal/logic"
	"chess-game/internal/protocol"
	"chess-game/model"
	"chess-game/pkg/format"
	"fmt"
)

func HandlerPlayerMove(game *model.Game, msg model.ClientMessage) {
	// Validate move
	err := logic.Moves(msg, game)
	if err != nil {
		fmt.Println("Movimento invalido ")

		fmt.Println(msg)

		message := fmt.Sprintf("🟢 It's still your turn ( %s )", err.Error())
		protocol.SendMessage(game.CurrentPlayer.Client, "TURN", message, true, format.Game(game))

		return
	}

	// Update board game
	game.Board = g.UpdateBoard(game.Chess.Position().Board().String())

	// Stop clock
	logic.StopClock(game.CurrentPlayer)

	// Change current player
	if game.Chess.Position().Turn() == 1 {
		game.CurrentPlayer = game.Players[0]
	} else if game.Chess.Position().Turn() == 2 {
		game.CurrentPlayer = game.Players[1]
	}

	// Start new clock
	logic.StartClock(game)

	// Update all clients
	protocol.SendMessage(game.CurrentPlayer.Client, "TURN", "🟢 It's your turn", true, format.Game(game))
	for _, player := range game.Players {
		if player != game.CurrentPlayer {
			protocol.SendMessage(player.Client, "WAIT", "🔴 Waiting for player to move", false, format.Game(game))
		}
	}
}
