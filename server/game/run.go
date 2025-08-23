package game

import (
	"chess-game/model"
	"fmt"

	"github.com/corentings/chess/v2"
)

var timeChan chan struct{}

func UpdatePlayers(player, oponnent *model.Player, game *model.Game) {
	GameFEN := game.Chess.Position().Board().String()

	msg := model.Message{
		Type: "TURN",
		Data: model.Data{
			White: model.PlayerFormat{
				Color:    game.Players[0].Color,
				Name:     game.Players[0].Name,
				Timeleft: game.Players[0].Timeleft,
			},
			Black: model.PlayerFormat{
				Color:    game.Players[1].Color,
				Timeleft: game.Players[1].Timeleft,
				Name:     game.Players[1].Name,
			},
			FEN:      GameFEN,
			Status:   fmt.Sprintf("%s turn", ColorName(player.Color)),
			LastMove: ReturnLastMove(game),
		},
	}

	oponnentMsg := msg
	oponnentMsg.Type = "WAIT"
	oponnent.Encoder.Encode(oponnentMsg)

	player.Encoder.Encode(msg)
}

func Run(game *model.Game) {
	defer func() {
		game.Players[0].Conn.Close()
		game.Players[1].Conn.Close()
	}()

	turn := 0
	for {
		currentPlayer := game.Players[turn]
		oponnent := game.Players[1-turn]

		UpdatePlayers(currentPlayer, oponnent, game)

		timeChan = make(chan struct{})
		go StartClock(currentPlayer, timeChan)

		moveChan := make(chan model.Message, 1)

		done := make(chan struct{})

		go func() {
			ok := ReadPlayerMessage(currentPlayer, moveChan)
			if !ok {
				EndGameByResign(currentPlayer, oponnent, game)
				close(done)
			}
		}()

		select {
		case <-done:
			return
		case <-timeChan:
			EndGameByTimeUp(currentPlayer, oponnent, game)
			return

		case playerMessage, ok := <-moveChan:
			if !ok {
				return
			}

			moveStr := playerMessage.Data.Message
			if IsValid, err := ValidMove(moveStr, &game.Chess); !IsValid {
				currentPlayer.Encoder.Encode(model.Message{
					Type: "INFO",
					Data: model.Data{Message: fmt.Sprintf("Invalid Move (%s)", err.Error())},
				})
				continue
			}

			if game.Chess.Outcome() != chess.NoOutcome {
				EndGame(currentPlayer, oponnent, game)
				return
			}

			close(timeChan)

			turn = 1 - turn
		}

	}
}
