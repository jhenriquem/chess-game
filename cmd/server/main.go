package main

import (
	"chess-game_cli/internal/server/core"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	addr := core.Run()

	flag.Parse()
	log.SetFlags(0)

	fmt.Println("Server online")
	log.Fatal(http.ListenAndServe(*addr, nil))
}
