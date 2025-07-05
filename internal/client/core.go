package client

import (
	"bufio"
	"chess-game/internal/model"
	"chess-game/internal/pkg/utils"
	"chess-game/internal/protocol"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func QuestionGame() bool {
	var wantPlay string
	fmt.Println("You are play? [Y/N]")
	_, err := fmt.Scanln(&wantPlay)
	if err != nil {
		return false
	}

	switch wantPlay {
	case "Y":
		return true
	default:
		return false

	}
}

var (
	Player = model.Player{}
	IsGame = false
)

func Run() {
	url := fmt.Sprintf("ws://%s/game", os.Getenv("SERVER"))

	utils.Introdution()
	continueGame := QuestionGame()
	if !continueGame {
		fmt.Println("👋 Bye")
		return
	}

	fmt.Println("Conectando ao servidor:", url)
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Erro ao conectar:", err)
	}

	Player.Client = conn

	if IsGame {
		SetPinglogic(conn)
	}

	defer conn.Close()

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

			if res.TypeInfo == "initGame" {
				IsGame = true
			}

			if IsGame {
				RenderBoard(res.Game.Board)

				fmt.Print("What's your move? :")

				scanner := bufio.NewScanner(os.Stdin)

				scanner.Scan()

				move := scanner.Text()
				SendMove(move, conn)
			}
		}
	}()

	select {
	case <-interrupt:
		log.Println("Encerrando conexão.")
		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		return
	}
}
