package server

import (
	g "chess-game/internal/game"
	"fmt"

	"github.com/corentings/chess/v2"
	"github.com/gorilla/websocket"
)

func HandleMatch(p1, p2 *websocket.Conn) {
	game := chess.NewGame()

	defer func() {
		p1.Close()
		p2.Close()
	}()

	gameId := g.New(p1, p2)

	for _, p := range g.GetOne(gameId).Players {
		go g.Monitoring(p, game)
	}

	fmt.Println(game.Position().Board())
	g.Run(gameId)

	// ... lógica da partida ...
}
