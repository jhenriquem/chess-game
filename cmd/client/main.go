package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".client.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Conectar ao servidor WebSocket
	url := fmt.Sprintf("ws://%s/game", os.Getenv("SERVER"))
	fmt.Println("Conectando ao servidor:", url)

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Erro ao conectar:", err)
	}
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(15 * time.Second))
	conn.SetPongHandler(func(appData string) error {
		conn.SetReadDeadline(time.Now().Add(15 * time.Second)) // renova deadline
		return nil
	})

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
