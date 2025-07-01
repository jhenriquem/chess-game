package server

import (
	g "chess-game_server/game"

	"github.com/gorilla/websocket"
)

func HandleMatch(p1, p2 *websocket.Conn) {
	defer func() {
		p1.Close()
		p2.Close()
	}()

	game := g.New(p1, p2)

	for _, p := range game.Players {
		go p.MonitoringConnection()
	}

	game.Run()

	// ... lógica da partida ...
}
