package client

import (
	"chess-game/internal/net"
	"chess-game/internal/protocol"
	"chess-game/internal/ui"
	"chess-game/model"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

func Run(url string) {
	player := model.PlayerFormat{}

	for {
		ui.ClearScreen()
		player.Name = AskName()
		if player.Name != "" {
			break
		}
		fmt.Println("Nome inválido")
	}

	ui.ClearScreen()

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Erro ao conectar:", err)
	}
	defer conn.Close()

	fmt.Println(" Conectado. Aguardando oponente...")

	inGame := false
	done := make(chan struct{})
	interrupt := make(chan os.Signal, 1)
	channel := make(chan protocol.Message)

	signal.Notify(interrupt, os.Interrupt)

	// Recebe mensagens do servidor
	go net.ReaderServer(conn, done, channel)

	go func() {
		for {
			select {
			case data := <-channel:
				if data.TypeInfo == "INIT" {
					idx := 1
					if data.IsTurn {
						idx = 0
					}
					player = data.Game.Players[idx]
				}

				for _, p := range data.Game.Players {
					if player.Color == p.Color {
						player = p
					}
				}

				ui.Load(data, &player)

				if !inGame {
					net.SetPingHandler(conn)
					inGame = true
				}

			case <-done:
				fmt.Println("Conexão encerrada pelo servidor.")
				return
			}
		}
	}()

	// Lê entradas do usuário (jogadas)
	go net.ClientInputLoop(conn, done)

	select {
	case <-interrupt:
		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Encerrado pelo cliente"))
	case <-done:
	}
}
