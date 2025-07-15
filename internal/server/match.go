package server

import (
	g "chess-game/internal/game"
	"chess-game/internal/net"

	"github.com/gorilla/websocket"
)

func HandleMatch(p1, p2 *websocket.Conn) {
	defer func() {
		p1.Close()
		p2.Close()
	}()

	game := g.New(p1, p2)

	for _, p := range game.Players {
		go func() {
			done := net.MonitoringClient(p.Client, p.Game)

			select {
			case <-done:
				p.Game.Desconnect <- p
				p.Client.Close()
			}
		}()
	}

	// fmt.Println(g.GetOne(gameId).Chess.Position().Board().Draw())

	g.Run(game)

	// ... lógica da partida ...
}
