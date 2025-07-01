package server

import (
	g "chess-game_server/game"
	"fmt"

	"github.com/gorilla/websocket"
)

func HandleMatch(p1, p2 *websocket.Conn) {
	game := g.New(p1, p2)

	go game.Run()

	go game.BlackPlayer.MonitoringConnection()
	go game.WhitePlayer.MonitoringConnection()

	fmt.Println("Create new room")

	// ... lógica da partida ...

	// p1.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Bye"))
	// p2.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Bye"))
}
