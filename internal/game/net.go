package game

import (
	"chess-game/internal/model"
	"chess-game/internal/pkg/format"
	"chess-game/internal/protocol"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

const (
	pongWait   = 30 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

func Monitoring(p *model.Player) {
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
		var clientMessage model.ClientMessage
		_, msg, err := p.Client.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			fmt.Println("Error reading player message :", err.Error())
			break

		}

		json.Unmarshal(msg, &clientMessage)

		fmt.Printf("%s player ( message type => %s)\n", p.Color, clientMessage.Type)
		fmt.Printf("Move %s player : %s\n\n", p.Color, strings.Join(clientMessage.Move, ""))

		protocol.SendMessage(p.Client, "playerMove", "Sua vez", format.ToFormatGame(p.Game))
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
