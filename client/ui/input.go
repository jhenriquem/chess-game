package ui

import (
	"github.com/gdamore/tcell/v2"
)

func Input(text string) {
	stText := tcell.StyleDefault.Background(tcell.ColorNone).Foreground(tcell.ColorWhite)

	screen.SetContent(6, 23, tcell.RuneULCorner, nil, stText)
	for x := 0; x < 10; x++ {
		screen.SetContent(7+x, 23, tcell.RuneHLine, nil, stText)
	}

	screen.SetContent(6, 24, tcell.RuneVLine, nil, stText)
	screen.SetContent(8, 24, '>', nil, stText)
	screen.SetContent(9, 24, ' ', nil, stText)

	for x, char := range text {
		screen.SetContent(10+x, 25, char, nil, stText)
	}

	screen.Show()
}
