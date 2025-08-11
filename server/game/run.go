package game

import (
	"chess-game/model"
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
		oponnent.Encoder.Encode(model.Message{Type: "WAIT", Data: model.Data{Message: "Waiting your opponent move"}})

		msg := model.Message{
			Type: "TURN",
			Data: model.Data{
				Player: model.PlayerFormat{
					Color:    currentPlayer.Color,
					Name:     currentPlayer.Name,
					Timeleft: currentPlayer.Timeleft,
				},
				Oponnent: model.PlayerFormat{
					Color:    oponnent.Color,
					Timeleft: oponnent.Timeleft,
					Name:     oponnent.Name,
				},
				FEN:     GameFEN,
				Message: "Is your turn",
			},
		}
		currentPlayer.Encoder.Encode(msg)

		// Start the timer
		timeChan = make(chan struct{})
		go StartClock(currentPlayer, timeChan)

		select {
		case <-timeChan:
			currentPlayer.Encoder.Encode(model.Message{Type: "INFO", Data: model.Data{Message: "Your time is up"}})
			oponnent.Encoder.Encode(model.Message{Type: "INFO", Data: model.Data{Message: "Your opponent's time is up"}})

			return
		default:
			defer close(timeChan)

			var currPlayerMove model.Message
			currentPlayer.Conn.SetReadDeadline(time.Now().Add(300 * time.Second))
			if err := currentPlayer.Decoder.Decode(&currPlayerMove); err != nil {
				log.Printf("\nError reading player(%s) message (%s)", currentPlayer.Color, err.Error())

				oponnent.Encoder.Encode(model.Message{Type: "INFO", Data: model.Data{Message: "Your opponent is desconnected"}})

				return
			}

			// Validação de movimento
			err := game.Chess.PushNotationMove(currPlayerMove.Data.Message, chess.UCINotation{}, nil)
			if err != nil {
				currentPlayer.Encoder.Encode(model.Message{Type: "TURN", Data: model.Data{Message: "Invalid Move"}})
				continue
			}

			// Atualização do turno
			turn = 1 - turn
			// Atualizar os players (tabulerio, tempo)
			// Trocar o player
		}

	}
}
