package game

import (
	"chess-game/model"
	"fmt"
	"log"
	"time"

	"github.com/corentings/chess/v2"
)

var timeChan chan struct{}

func Run(game *model.Game) {
	defer func() {
		game.Players[0].Conn.Close()
		game.Players[1].Conn.Close()
	}()

	turn := 0
	for {
		GameFEN := game.Chess.Position().Board().String()

		currentPlayer := game.Players[turn]
		oponnent := game.Players[1-turn]

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
				Status:   fmt.Sprintf("%s turn", ColorName(currentPlayer.Color)),
				LastMove: ReturnLastMove(game),
			},
		}

		oponnentMsg := msg
		oponnentMsg.Type = "WAIT"
		oponnent.Encoder.Encode(oponnentMsg)

		currentPlayer.Encoder.Encode(msg)

		timeChan = make(chan struct{})
		go StartClock(currentPlayer, timeChan)

		moveChan := make(chan model.Message, 1)
		go func() {
			ReadPlayerMessage(currentPlayer, oponnent, moveChan)
		}()

		select {
		case <-timeChan:
			currentPlayer.Encoder.Encode(model.Message{Type: "INFO", Data: model.Data{Message: "Your time is up"}})
			oponnent.Encoder.Encode(model.Message{Type: "INFO", Data: model.Data{Message: "Your opponent's time is up"}})
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
				result := game.Chess.Outcome().String()
				currentPlayer.Encoder.Encode(model.Message{Type: "INFO", Data: model.Data{Message: "Game Over: " + result}})
				oponnent.Encoder.Encode(model.Message{Type: "INFO", Data: model.Data{Message: "Game Over: " + result}})
				return
			}

			close(timeChan)

			turn = 1 - turn
		}

	}
}

func ReadPlayerMessage(player, oponnent *model.Player, moveChan chan model.Message) {
	var playerMessage model.Message

	player.Conn.SetReadDeadline(time.Now().Add(300 * time.Second))

	if err := player.Decoder.Decode(&playerMessage); err != nil {
		log.Printf("\nError reading player(%s) message (%s)", player.Color, err.Error())

		oponnent.Encoder.Encode(model.Message{
			Type: "INFO",
			Data: model.Data{Message: "Your opponent is disconnected"},
		})

		close(moveChan)

		return
	}
	moveChan <- playerMessage
}

func ValidMove(move string, c *chess.Game) (bool, error) {
	for _, validMove := range c.Position().ValidMoves() {
		if validMove.String() == move {
			err := c.PushNotationMove(move, chess.UCINotation{}, nil)
			if err != nil {
				return false, err
			}
			return true, nil
		}
	}
	return false, fmt.Errorf("Invalid Move")
}
