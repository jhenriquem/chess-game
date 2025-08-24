package main

import (
	"chess-game/server/handler"
	"flag"
	"log"
	"net/http"
	"os"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	http.HandleFunc("/game", handler.Game)

	log.Println("Server listening on :8000")

	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
}
