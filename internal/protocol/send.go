package protocol

import (
	"chess-game/internal/model"
	"fmt"

	"github.com/gorilla/websocket"
)

func SendMessage(conn *websocket.Conn, typeInfo, info string, turn bool, game model.Protogame) error {
	body := Message{
		TypeInfo: typeInfo,
		Info:     info,
		Game:     game,
		IsTurn:   turn,
	}

	err := conn.WriteJSON(body)
	if err != nil {
		return fmt.Errorf("Error sending message : %v", err.Error())
	}
	return nil
}
