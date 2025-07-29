package client

import (
	"chess-game/internal/net"
	"chess-game/internal/protocol"
	"chess-game/internal/ui"
	"chess-game/model"
	"chess-game/pkg/utils"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

var (
	inGame = false
	Player = model.PlayerFormat{}
)

func Run(url string) {
	utils.Introdution()
	continueGame := utils.QuestionGame()
	if !continueGame {
		fmt.Println("👋 Bye")
		return
	}

	ui.ClearScreen()

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Erro ao conectar:", err)
	}
	defer conn.Close()

	done := make(chan struct{})
	interrupt := make(chan os.Signal, 1)
	channel := make(chan protocol.Message)

	signal.Notify(interrupt, os.Interrupt)

	go net.ReaderServer(conn, done, channel)
	go func() {
		for {
			select {
			case data := <-channel:
				if data.TypeInfo == "INIT" && data.IsTurn {
					Player = data.Game.Players[0]
				} else if data.TypeInfo == "INIT" && !data.IsTurn {
					Player = data.Game.Players[1]
				}

				// Constant update of the Player
				for _, player := range data.Game.Players {
					if Player.Color == player.Color {
						Player = player
					}
				}

				ui.Load(data, &Player)

				if !inGame {
					net.SetPingHandler(conn) // Ativa o monitoramento de ping
					inGame = true
				}
			case <-done:
				return
			}
		}
	}()

	go net.ClientInputLoop(conn, done)

	select {
	case <-interrupt:
		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	case <-done:
		log.Println("Conexão encerrada pelo servidor.")
	}
}
