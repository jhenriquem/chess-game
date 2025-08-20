package ui

import (
	"chess-game/model"
	"fmt"

	"github.com/gdamore/tcell/v2"
)

func Header(data model.Data) {
	stText := tcell.StyleDefault.Background(tcell.ColorNone).Foreground(tcell.ColorWhite)

	playerSide := fmt.Sprintf("%s (%s) White", data.White.Name, data.White.Timeleft)
	oponnentSide := fmt.Sprintf("Black %s (%s)", data.Black.Name, data.Black.Timeleft)

	header := fmt.Sprintf("%s X %s", playerSide, oponnentSide)

	for x := 0; x < 60; x++ {
		screen.SetContent(2+x, 1, tcell.RuneHLine, nil, stText)
	}

	for x, char := range header {
		screen.SetContent(2+x, 2, char, nil, stText)
	}

	for x := 0; x < 60; x++ {
		screen.SetContent(2+x, 3, tcell.RuneHLine, nil, stText)
	}

	screen.Show()
}
