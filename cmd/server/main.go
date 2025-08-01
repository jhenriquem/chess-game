package main

import (
	"chess-game/internal/server"
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

	addr := server.Run()

	fmt.Printf("Server online -> %s\n ", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
