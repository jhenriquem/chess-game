package main

import (
	"chess-game/server/handler"
	"log"
	"net"
	"sync"
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	log.Println("Server listening on :8000")

	var wg sync.WaitGroup

	// accept loop
	wg.Add(1)

	handler.Connection(ln)

	wg.Wait()
}
