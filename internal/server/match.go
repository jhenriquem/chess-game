package server

import (
	g "chess-game/internal/game"
	"chess-game/internal/handler"
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
			message := make(chan model.ClientMessage)

			go handler.HandleClientMessage(game, message, done)
			net.MonitoringClient(p.Client, p.Game, message, done)

			<-done
			game.Desconnect <- p
		}(player)
	}

	g.StartGame(game)
}
