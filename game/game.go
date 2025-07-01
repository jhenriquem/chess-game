package game

import (
	"fmt"
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
		p.SendInfo(msg)
	}

	for {
		select {
		case player := <-g.Desconnect:
			msg := "Outher player desconnected, you win"
			if player.Color == "black" {
				g.GetWhite().SendInfo(msg)
				g.GetWhite().Client.Close()
			} else {
				g.GetBlack().SendInfo(msg)
				g.GetBlack().Client.Close()
			}

			break
		}
	}
}

func (g *Game) SetHouse(positon [2]int, piece *Pieces) {
	g.Board[positon[1]][positon[0]] = piece
}
