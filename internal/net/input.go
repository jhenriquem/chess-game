package net

import (
	"bufio"
	"chess-game/model"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func ClientInputLoop(conn *websocket.Conn, done <-chan struct{}) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		select {
		case <-done:
			return
		default:
			if !scanner.Scan() {
				log.Println("Entrada encerrada.")
				return
			}

			move := scanner.Text()

			msg := model.ClientMessage{
				Type: "MOVE",
				Move: move,
			}

			if err := conn.WriteJSON(msg); err != nil {
				log.Println("Erro ao enviar mensagem:", err)
				return
			}
		}
	}
}
