package game

import (
	"chess-game/pkg/protocol"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func (g *Game) ToFormatForProtocol() protocol.Game {
	return protocol.Game{
		Board:   g.ToFormatBoard(),
		Moves:   g.Moves,
		Players: [2]protocol.Player{g.Players[0].ToFormatForJson(), g.Players[1].ToFormatForJson()},
		Timer:   g.Timer,
		Turn:    g.Turn.Color,
	}
}

const (
	pongWait   = 30 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

func (p *Player) MonitoringConnection() {
	defer func() {
		p.Game.Desconnect <- p
		p.Client.Close()
	}()

	p.Client.SetReadDeadline(time.Now().Add(pongWait))
	p.Client.SetPongHandler(func(string) error {
		p.Client.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	done := make(chan struct{})

	go pingLoop(p.Client, done)

	for {
		_, _, err := p.Client.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
	}

	close(done) // avisa para parar os pings
}

func pingLoop(client *websocket.Conn, done <-chan struct{}) {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			if err := client.WriteControl(websocket.PingMessage, nil, time.Now().Add(pongWait)); err != nil {
				log.Println("Erro ao enviar ping:", err)
				return
			}
		}
	}
}
