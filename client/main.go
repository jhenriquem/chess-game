package main

import (
	"chess-game/client/net"
	"chess-game/client/ui"
	"chess-game/model"
	"fmt"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

func main() {
	s := ui.InitScreen()

	defer s.Fini()

	if len(os.Args) < 2 {
		fmt.Println("Uso: client <name>")
		os.Exit(1)
	}

	name := os.Args[1]

	client, err := net.ConnectedServer(name)
	if err != nil {
		log.Fatalf("Failed to connect to server (%s)", err.Error())
	}

	defer client.Conn.Close()

	messageChan := make(chan model.Message)
	eventChan := make(chan tcell.Event)
	errChan := make(chan error)

	go client.ReadServer(messageChan, errChan)
	go PollEventLoop(s, eventChan)

	GameLoop(client, s, messageChan, eventChan, errChan)
}

func PollEventLoop(s tcell.Screen, eventChan chan tcell.Event) {
	for {
		ev := s.PollEvent()
		eventChan <- ev
	}
}

func GameLoop(client *net.Client, s tcell.Screen, messageChan <-chan model.Message, eventChan <-chan tcell.Event, errChan <-chan error) {
	running := true

	input := ""

	for running {
		select {
		case message := <-messageChan:

			s.Clear()

			if message.Data.FEN != "" {
				ui.RenderBoard(message.Data.FEN)
			}

			ui.StatusBar(message.Data)

		case err := <-errChan:
			log.Printf("Erro de conexão: %s", err)
			running = false

		case ev := <-eventChan:
			switch ev := ev.(type) {

			case *tcell.EventResize:
				s.Sync()

			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					running = false
					return

				} else if ev.Key() == tcell.KeyEnter {

					client.SendMove(input)
					input = ""

				} else {

					input += string(ev.Rune())
					ui.Input(input)

				}
			}
		}
	}
}
