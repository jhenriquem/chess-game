package net

import (
	"bufio"
	"chess-game/internal/model"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func ClientInputLoop(conn *websocket.Conn, done <-chan struct{}) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		select {
		case <-done:
			break // encerra quando a conexão é finalizada
		default:
			if !scanner.Scan() {
				log.Println("Entrada encerrada.")
				return
			}

			text := scanner.Text()

			moves := []string{}

			for _, char := range text {
				moves = append(moves, string(char))
			}

			msg := model.ClientMessage{
				Type: "move",
				Move: moves,
			}

			if err := conn.WriteJSON(msg); err != nil {
				log.Println("Erro ao enviar mensagem:", err)
				return
			}
		}
	}
}
