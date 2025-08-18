package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	b := widgets.NewParagraph()
	b.Text = "Hello World!"
	b.SetRect(25, 5, 50, 10)

	ui.Render(b)

	text := ""
	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID { // event string/identifier
			case "<C-c>": // press 'q' or 'C-c' to quit
				return
			case "<MouseLeft>":
			case "<Resize>":
			default:

				text += e.ID

				b.Text = text
				ui.Render(b)
			}
			switch e.Type {
			case ui.KeyboardEvent: // handle all key presses
			}
		}
	}
}
