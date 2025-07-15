package client

import (
	"chess-game/internal/model"
	"chess-game/internal/net"
	"chess-game/internal/pkg/utils"
	"chess-game/internal/protocol"
	"chess-game/internal/ui"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

var (
	inGame = false
	Player = model.Protoplayer{}
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
				if data.Info != "" {
					fmt.Println("🟢", data.Info)
				}

				if data.Game.Turn != "" {
					ui.Load(data.Game.Board, true)
				}

				if !inGame {
					net.SetPingHandler(conn) // Ativa o monitoramento de ping
					inGame = true
				}
			case <-done:
				break
			}
		}
	}()

	go net.ClientInputLoop(conn, done)

	select {
	case <-interrupt:
		log.Println("Interrompido. Encerrando conexão.")
		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	case <-done:
		log.Println("Conexão encerrada pelo servidor.")

	}
}
