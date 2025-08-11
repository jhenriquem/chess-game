package main

import (
	"chess-game/model"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: client <name>")
		os.Exit(1)
	}

	name := os.Args[1]

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	enc := json.NewEncoder(conn)
	dec := json.NewDecoder(conn)

	msg := model.Message{
		Type: "CONNECTED",
		Data: model.Data{
			FEN:     "",
			Message: name,
		},
	}
	enc.Encode(msg)

	for {
		var m model.Message
		if err := dec.Decode(&m); err != nil {
			log.Printf("\nDisconnect or decode error (%s)", err.Error())
			return
		}

		fmt.Printf("\n FEN : %s", m.Data.FEN)
		fmt.Printf("\n Messsage : %s", m.Data.Message)

	}
}
