package game

import (
	"chess-game/internal/logic"
	"chess-game/internal/protocol"
	"chess-game/model"
	"chess-game/pkg/format"
)

func Start(game *model.Game) {
	// White Player message
	protocol.SendMessage(game.CurrentPlayer.Client, "INIT", "📌 You are playing, you are white ⬜. \n🟢 It's your turn  ", true, format.Game(game))

	// Black Player message
	msg := "📌 You are playing, you are black ⬛"
	protocol.SendMessage(game.GetPlayer("B").Client, "INIT", msg, false, format.Game(game))

	// Start time clock of CurrentPlayer
	logic.StartClock(game)
}
