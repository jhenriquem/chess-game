package server

import (
	g "chess-game/internal/game"
	"chess-game/internal/net"
	"chess-game/model"

	"github.com/gorilla/websocket"
)

func HandleMatch(p1, p2 *websocket.Conn) {
	game := g.New(p1, p2)

	// Inicia monitoramento para cada jogador
	for _, player := range game.Players {
		go func(p *model.Player) {
			done := make(chan struct{})
			net.MonitoringClient(p.Client, p.Game, done)

			<-done
			game.Desconnect <- p
		}(player)
	}

	g.Run(game)
}
