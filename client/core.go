package client

import (
	"chess-game/pkg/protocol"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func Run() {
	url := fmt.Sprintf("ws://%s/game", os.Getenv("SERVER"))
	fmt.Println("Conectando ao servidor:", url)

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Erro ao conectar:", err)
	}

	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	conn.SetPingHandler(func(appData string) error {
		conn.SetReadDeadline(time.Now().Add(30 * time.Second))
		conn.WriteControl(websocket.PongMessage, nil, time.Now().Add(30*time.Second))
		return nil
	})

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Erro ao ler mensagem:", err)
				return
			}

			var res protocol.Message
			if err := json.Unmarshal(msg, &res); err != nil {
				fmt.Println("Mensagem recebida (bruta):", string(msg))
				log.Println("Erro ao decodificar JSON:", err)
				continue
			}

			clearScreen()
			if res.Info != "" {
				color.Green("Servidor: %s", res.Info)
			}
			RenderBoard(res.Game.Board)
		}
	}()

	select {
	case <-interrupt:
		log.Println("Encerrando conexão.")
		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		return
	}
}
