package game

import (
	"chess-game/model"
	"time"
)

func StartClock(player *model.Player, timeChan chan struct{}) {
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				player.Timeleft -= 1 * time.Second

				// fmt.Printf("\n Timer clock of player (%s) : %v", player.Color, player.Timeleft)

				if player.Timeleft <= 0 {
					player.Timeleft = 0

					close(timeChan)

					return
				}
			case <-timeChan:
				return
			}
		}
	}()
}
