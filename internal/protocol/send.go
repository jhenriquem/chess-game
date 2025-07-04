package protocol

import (
	"chess-game/internal/model"
	"fmt"

	"github.com/gorilla/websocket"
)

func SendMessage(conn *websocket.Conn, info string, game model.Protogame) error {
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
