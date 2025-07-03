package protocol

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func SendMessage(conn *websocket.Conn, info string, game Game) error {
	body := Message{
		Info: info,
		Game: game,
	}

	err := conn.WriteJSON(body)
	if err != nil {
		return fmt.Errorf("Error sending message : %v", err.Error())
	}
	return nil
}
