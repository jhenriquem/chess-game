package logic

import (
	"chess-game/model"
	"time"
)

func StartClock(p *model.Player) {
	p.StopChan = make(chan struct{})
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				p.TimeLeft -= 1 * time.Second
				if p.TimeLeft <= 0 {
					p.TimeLeft = 0
					p.Game.Desconnect <- p
					return
				}
			case <-p.StopChan:
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
