package client

import (
	"chess-game/internal/net"
	"chess-game/internal/protocol"
	"chess-game/internal/ui"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func ReadServer(conn *websocket.Conn, done chan struct{}) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Erro ao ler mensagem:", err)
			close(done)
			break
		}

		var data protocol.Message
		if err := json.Unmarshal(msg, &data); err != nil {
			log.Println("Erro ao decodificar JSON:", err)
			fmt.Println("Mensagem bruta:", string(msg))
			continue
		}

		fmt.Println()
		if data.Info != "" {
			fmt.Println("🟢", data.Info)
		}

		if data.Game.Turn != "" {

			isTurn := false

			ui.Load(data.Game.Board, isTurn)

			if !inGame {
				net.SetPingLogic(conn) // Ativa o monitoramento de ping
				go net.StartPinger(conn, done)

				inGame = true
			}
		}
	}
}
