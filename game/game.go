package game

import (
	"chess-game/pkg/protocol"
	"fmt"

	"github.com/gorilla/websocket"
)

func (g *Game) GetBlack() *Player {
	if g.Players[0].Color == "black" {
		return g.Players[0]
	} else {
		return g.Players[1]
	}
}

func (g *Game) GetWhite() *Player {
	if g.Players[0].Color == "white" {
		return g.Players[0]
	} else {
		return g.Players[1]
	}
}

func (g *Game) Run() {
	for _, p := range g.Players {
		msg := fmt.Sprintf("You are playing, you are %s", p.Color)
		protocol.SendMessage(p.Client, msg, g.ToFormatForProtocol())
	}

	for {
		select {
		case player := <-g.Desconnect:
			msg := "Outher player desconnected, you win"
			if player.Color == "black" {
				protocol.SendMessage(g.GetWhite().Client, msg, g.ToFormatForProtocol())

				g.GetWhite().Client.WriteMessage(websocket.CloseMessage, []byte{})
				g.GetWhite().Client.Close()
			} else {
				protocol.SendMessage(g.GetBlack().Client, msg, g.ToFormatForProtocol())
				g.GetBlack().Client.WriteMessage(websocket.CloseMessage, []byte{})
				g.GetBlack().Client.Close()
			}

			break
		}
	}
}

func (g *Game) SetHouse(positon [2]int, piece *Pieces) {
	g.Board[positon[1]][positon[0]] = piece
}
