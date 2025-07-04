package server

import (
	g "chess-game/internal/game"

	"github.com/gorilla/websocket"
)

func HandleMatch(p1, p2 *websocket.Conn) {
	defer func() {
		p1.Close()
		p2.Close()
	}()

	gameId := g.New(p1, p2)

	for _, p := range g.GetOne(gameId).Players {
		go g.MonitoringConnection(p)
	}

	g.Run(gameId)

	// ... lógica da partida ...
}
