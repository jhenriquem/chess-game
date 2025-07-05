package client

import (
	"chess-game/internal/model"
	"fmt"

	"github.com/gorilla/websocket"
)

func SendMove(move string, conn *websocket.Conn) {
	moves := []string{}

	for _, char := range move {
		moves = append(moves, string(char))
	}

	msg := model.ClientMessage{
		Type: "move",
		Move: moves,
	}

	if err := conn.WriteJSON(msg); err != nil {
		fmt.Println("Error : ", err.Error())
	}
	fmt.Println("skdjnskdjn")
}
