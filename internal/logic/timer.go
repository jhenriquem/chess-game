package logic

import (
	"chess-game/model"
	"fmt"
	"time"
)

func StartClock(g *model.Game) {
	g.CurrentPlayer.StopChan = make(chan struct{})
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				g.CurrentPlayer.TimeLeft -= 1 * time.Second

				fmt.Printf("\n Timer clock of player (%s) : %v", g.CurrentPlayer.Color, g.CurrentPlayer.TimeLeft)

				if g.CurrentPlayer.TimeLeft <= 0 {
					g.CurrentPlayer.TimeLeft = 0
					g.CurrentPlayer.Game.Timeout <- g.CurrentPlayer
					return
				}
			case <-g.CurrentPlayer.StopChan:
				return
			}
		}
	}()
}

func StopClock(p *model.Player) {
	if p.StopChan != nil {
		close(p.StopChan)
	}
}
