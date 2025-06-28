package room

import (
	"chess-game_cli/internal/game"
	"chess-game_cli/pkg/model"

	"github.com/gorilla/websocket"
)

type Room struct {
	Game        *game.GameStruct
	BlackPlayer *model.Player
	WhitePlayer *model.Player
	Timer       string
	Turn        *model.Player
}

func newPlayer(client *websocket.Conn, color string) *model.Player {
	return &model.Player{
		Client:      client,
		ColorPieces: color,
	}
}

func NewRoom(p1, p2 *websocket.Conn) *Room {
	return &Room{
		Game:        game.New(),
		WhitePlayer: newPlayer(p1, "white"),
		BlackPlayer: newPlayer(p2, "black"),
	}
}
