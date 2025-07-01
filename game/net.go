package game

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	pongWait   = 10 * time.Second // tempo para esperar o pong
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
			log.Println("Erro na leitura:", err)
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
			if err := client.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Println("Erro ao enviar ping:", err)
				return
			}
		}
	}
}
