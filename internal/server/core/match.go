package core

import (
	"chess-game_cli/internal/server/room"
	"fmt"

	"github.com/gorilla/websocket"
)

func HandleMatch(p1, p2 *websocket.Conn) {
	defer p1.Close()
	defer p2.Close()

	r := room.NewRoom(p1, p2)
	r.Game.InitBoard(r.WhitePlayer, r.BlackPlayer)

	fmt.Println("Create new room")
	// ... lógica da partida ...

	p1.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Bye"))
	p2.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Bye"))
}
