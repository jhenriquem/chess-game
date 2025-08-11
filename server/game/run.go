package game

import (
	"chess-game/model"
	"fmt"
	"log"
	"time"
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
				FEN:     GameFEN,
				Message: fmt.Sprintf("Is your turn"),
			},
		}
		currentPlayer.Encoder.Encode(msg)

		// Start the timer
		timeChan = make(chan struct{})
		go StartClock(currentPlayer, timeChan)

		select {
		case <-timeChan:
			currentPlayer.Encoder.Encode(model.Message{Type: "INFO", Data: model.Data{FEN: GameFEN, Message: "Your time is up"}})
			oponnent.Encoder.Encode(model.Message{Type: "INFO", Data: model.Data{FEN: GameFEN, Message: "Your opponent's time is up"}})

			return
		default:
			defer close(timeChan)

			var currPlayerMove model.Message
			currentPlayer.Conn.SetReadDeadline(time.Now().Add(300 * time.Second))
			if err := currentPlayer.Decoder.Decode(&currPlayerMove); err != nil {
				log.Printf("\nError reading player(%s) message (%s)", currentPlayer.Color, err.Error())

				oponnent.Encoder.Encode(model.Message{Type: "INFO", Data: model.Data{FEN: "", Message: "Your opponent is desconnected"}})

				return
				// oponnent.Encoder.Encode()
			}

			// Validação de movimento
			// Parar o cronometro
			// Atualizar os players (tabulerio, tempo)
			// Trocar o player
		}

	}
}
