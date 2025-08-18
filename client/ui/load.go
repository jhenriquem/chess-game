package ui

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

var screen tcell.Screen

func InitScreen() tcell.Screen {
	var err error
	screen, err = tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	return screen
}
