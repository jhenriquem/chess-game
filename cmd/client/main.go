package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	// Conectar ao servidor WebSocket
	url := "ws://localhost:8080/game"
	fmt.Println("Conectando ao servidor:", url)

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Erro ao conectar:", err)
	}
	defer conn.Close()

	// Canal para leitura de mensagens do servidor
	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Erro ao ler mensagem:", err)
				return
			}
			fmt.Println("Servidor:", string(msg))
		}
	}()

	// Leitura de mensagens do terminal para enviar ao servidor
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Você: ")
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()
		err := conn.WriteMessage(websocket.TextMessage, []byte(text))
		if err != nil {
			log.Println("Erro ao enviar:", err)
			return
		}
	}
}
