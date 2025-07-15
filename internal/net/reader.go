package net

import (
	"chess-game/internal/model"
	"chess-game/internal/protocol"
	"fmt"
	"log"
	"strings"

	"github.com/gorilla/websocket"
)

func ReaderServer(conn *websocket.Conn, done chan struct{}, message chan protocol.Message) {
	for {
		var data protocol.Message

		err := conn.ReadJSON(&data)
		if err != nil {
			log.Println("Erro ao ler mensagem:", err)
			close(done)
			break
		}

		message <- data
	}
}

func ReaderClient(conn *websocket.Conn, game *model.Game, message chan model.ClientMessage) {
	for {
		var clientMessage model.ClientMessage

		err := conn.ReadJSON(&clientMessage)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			fmt.Println("Error reading player message :", err.Error())
			break
		}

		fmt.Printf("Move player : %s\n\n", strings.Join(clientMessage.Move, ""))

		message <- clientMessage
	}
}
