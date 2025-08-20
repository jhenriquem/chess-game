package ui

import (
	"github.com/gdamore/tcell/v2"
)

func Input(text string) {
	stText := tcell.StyleDefault.Background(tcell.ColorNone).Foreground(tcell.ColorWhite)

	screen.SetContent(3, 19, '>', nil, stText)
	screen.SetContent(4, 19, ' ', nil, stText)

	for x, char := range text {
		screen.SetContent(5+x, 19, char, nil, stText)
	}

	screen.Show()
}
