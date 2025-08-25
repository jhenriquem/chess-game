package main

import (
	"chess-game/client/ui"
	"chess-game/model"
	"chess-game/net"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/corentings/chess/v2"
	"github.com/gdamore/tcell/v2"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
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

	var timeLeft time.Duration
	timeStop := make(chan struct{})
	var once sync.Once

	closeTime := func() {
		once.Do(func() {
			close(timeStop)
		})
	}

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

	startTimer := func(initial time.Duration) {
		timeLeft = initial
		go func() {
			ticker := time.NewTicker(1 * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-ticker.C:
					timeLeft -= time.Second

					switch color {
					case chess.Black:
						lastMessage.Data.Black.Timeleft = timeLeft
					case chess.White:
						lastMessage.Data.White.Timeleft = timeLeft
					}

					if timeLeft <= 0 {
						timeLeft = 0
						ui.Header(lastMessage.Data)
						closeTime()
						return
					}
					ui.Header(lastMessage.Data)
				case <-timeStop:
					return
				}
			}
		}()
	}

	for {
		select {
		case message := <-messageChan:
			lastMessage = message

			switch color {
			case chess.Black:
				timeLeft = lastMessage.Data.Black.Timeleft
			case chess.White:
				timeLeft = lastMessage.Data.White.Timeleft
			}

			if lastMessage.Type == "TURN" {
				timeStop = make(chan struct{})
				once = sync.Once{}
				startTimer(timeLeft)
			} else if lastMessage.Type == "WAIT" {
				closeTime()
			}

			if lastMessage.Type == "END" {
				closeTime()

				msg := lastMessage.Data.Message
				s.Clear()
				ui.Header(lastMessage.Data)
				ui.RenderBoard(lastMessage.Data.FEN, color)
				ui.StatusBar(lastMessage.Data)
				ui.PrintMessage(msg)

				for {
					ev := s.PollEvent()
					switch ev := ev.(type) {
					case *tcell.EventKey:
						if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyEnter {
							close(done)
							return
						}
					}
				}
			}

			render()

		case <-done:
			closeTime()
			return

		case err := <-errChan:
			msg := fmt.Sprintf("Error connecting to server (%s)", err.Error())
			ui.PrintMessage(msg)
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
