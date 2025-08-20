package main

import (
	"chess-game/client/net"
	"chess-game/client/ui"
	"chess-game/model"
	"fmt"
	"log"
	"os"

	"github.com/corentings/chess/v2"
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
		log.Fatalf("Falha ao conectar ao servidor: %s", err)
	}
	defer client.Conn.Close()

	messageChan := make(chan model.Message)
	eventChan := make(chan tcell.Event)
	errChan := make(chan error)
	done := make(chan struct{})

	go client.ReadServer(messageChan, errChan)
	go PollEventLoop(s, eventChan)

	GameLoop(client, s, messageChan, eventChan, errChan, done, name)
}

func PollEventLoop(s tcell.Screen, eventChan chan tcell.Event) {
	for {
		ev := s.PollEvent()
		eventChan <- ev
	}
}

func GameLoop(client *net.Client, s tcell.Screen, messageChan <-chan model.Message, eventChan <-chan tcell.Event, errChan <-chan error, done chan struct{}, name string) {
	input := ""
	var lastMessage model.Message
	color := chess.White

	render := func() {
		s.Clear()

		if lastMessage.Data.FEN != "" {

			if lastMessage.Data.Black.Name == name {
				color = chess.Black
			}
			ui.Header(lastMessage.Data)
			ui.RenderBoard(lastMessage.Data.FEN, color)
		}

		if lastMessage.Type == "TURN" {
			ui.Input(input)
		}
		ui.StatusBar(lastMessage.Data)
		s.Show()
	}

	for {
		select {
		case <-done:
			return

		case message := <-messageChan:
			lastMessage = message
			render()

		case err := <-errChan:
			log.Printf("Erro de conexão: %s", err)
			select {
			case <-done:
			default:
				close(done)
			}

		case ev := <-eventChan:
			switch ev := ev.(type) {
			case *tcell.EventResize:
				s.Sync()
				render()

			case *tcell.EventKey:
				var exit bool
				input, exit = KeyEvents(ev, input, client)
				if exit {
					select {
					case <-done:
					default:
						close(done)
					}
				}
				render()
			}
		}
	}
}

func KeyEvents(ev *tcell.EventKey, input string, client *net.Client) (string, bool) {
	switch ev.Key() {
	case tcell.KeyEscape, tcell.KeyCtrlC:
		return "", true

	case tcell.KeyBackspace, tcell.KeyBackspace2:
		if len(input) >= 1 {
			input = input[:len(input)-1]
		}
		return input, false

	case tcell.KeyEnter:
		if input != "" {
			client.SendMove(input)
			input = ""
		}
		return input, false

	default:
		if ev.Rune() != 0 {
			input += string(ev.Rune())
		}
	}
	return input, false
}
