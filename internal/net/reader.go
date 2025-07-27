package net

import (
	"chess-game/internal/protocol"
	"chess-game/model"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func ReaderServer(conn *websocket.Conn, done chan struct{}, message chan protocol.Message) {
	for {
		var data protocol.Message

		err := conn.ReadJSON(&data)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}

			log.Println("Erro ao ler mensagem:", err)
			close(done)
			break
		}

		message <- data
	}
}

func ReaderClient(conn *websocket.Conn, game *model.Game, message chan model.ClientMessage, done chan struct{}) {
	defer close(done)
	for {
		var clientMessage model.ClientMessage

		err := conn.ReadJSON(&clientMessage)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			fmt.Println("Error reading player message :", err.Error())
			return
		}
		message <- clientMessage
	}
}
