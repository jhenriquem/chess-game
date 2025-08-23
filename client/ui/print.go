package ui

import "github.com/gdamore/tcell/v2"

func PrintMessage(msg string) {
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	startX := 3
	startY := 19

	screen.SetContent(2, startY, ' ', nil, style)
	for i, r := range msg {
		screen.SetContent(startX+i, startY, r, nil, style)
	}
	screen.SetContent(startX+len(msg), startY, ' ', nil, style)

	screen.Show()
}
