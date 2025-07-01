package game

import (
	"github.com/gorilla/websocket"
)

func (g *Game) Run() {
	g.BlackPlayer.Client.WriteMessage(websocket.TextMessage, []byte("Your are playing"))
	g.WhitePlayer.Client.WriteMessage(websocket.TextMessage, []byte("Your are playing"))

	for {
		select {
		case player := <-g.Desconnect:
			if player.ColorPieces == "black" {
				g.WhitePlayer.Client.WriteMessage(websocket.TextMessage, []byte("Outher player desconnected"))
			} else {
				g.BlackPlayer.Client.WriteMessage(websocket.TextMessage, []byte("Outher player desconnected"))
			}
		}
	}
}

func (g *Game) SetHouse(positon [2]int, piece *Pieces) {
	g.Board[positon[1]][positon[0]] = piece
}
