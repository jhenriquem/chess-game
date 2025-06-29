package main

import (
	"chess-game_cli/internal/server/core"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".server.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	addr := core.Run()

	fmt.Printf("Server online -> %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
