package main

import (
	"chess-game/client"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".client.env")
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	client.Run()
}
