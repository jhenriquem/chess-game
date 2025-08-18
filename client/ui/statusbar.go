package ui

import (
	"chess-game/model"
	"fmt"

	"github.com/gdamore/tcell/v2"
)

func StatusBar(data model.Data) {
	stText := tcell.StyleDefault.Background(tcell.ColorNone).Foreground(tcell.ColorWhite)

	msg := data.Message

	if data.FEN == "" {
		str := fmt.Sprintf("Status : %s", msg)
		for x, char := range str {
			screen.SetContent(2+x, 2, char, nil, stText)
		}
		screen.Show()
		return
	}

	y := 14

	for x := 0; x < 40; x++ {
		screen.SetContent(2+x, y-1, tcell.RuneHLine, nil, stText)
	}

	str := fmt.Sprintf("Last move : %s", data.LastMove)
	for x, char := range str {
		screen.SetContent(2+x, y, char, nil, stText)
	}

	status := fmt.Sprintf("Status : %s", msg)
	for x, char := range status {
		screen.SetContent(10+len(str)+x, y, char, nil, stText)
	}

	screen.Show()
}
