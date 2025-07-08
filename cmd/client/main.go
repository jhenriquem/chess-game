package main

import (
	"chess-game/internal/client"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".client.env")
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	url := fmt.Sprintf("ws://%s/game", os.Getenv("SERVER"))

	client.Run(url)
}
