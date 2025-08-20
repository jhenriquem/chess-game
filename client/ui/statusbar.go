package ui

import (
	"chess-game/model"
	"fmt"

	"github.com/gdamore/tcell/v2"
)

func StatusBar(data model.Data) {
	stText := tcell.StyleDefault.Background(tcell.ColorNone).Foreground(tcell.ColorWhite)

	msg := data.Message
	statusMsg := data.Status

	if data.FEN == "" {
		str := fmt.Sprintf(" %s ", msg)
		for x, char := range str {
			screen.SetContent(2+x, 1, char, nil, stText)
		}
		screen.Show()
		return
	}

	y := 17

	for x := 0; x < 60; x++ {
		screen.SetContent(2+x, y-1, tcell.RuneHLine, nil, stText)
	}

	str := fmt.Sprintf("Last move : %s", data.LastMove)
	for x, char := range str {
		screen.SetContent(2+x, y, char, nil, stText)
	}

	status := fmt.Sprintf("Status : %s", statusMsg)
	for x, char := range status {
		screen.SetContent(10+len(str)+x, y, char, nil, stText)
	}

	screen.Show()
}
