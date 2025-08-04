package model

import "time"

type PlayerFormat struct {
	Color    string // W or B
	TimeLeft time.Duration
	Moves    []string
	Score    int
	Name     string
}

type GameFormat struct {
	Board [8][8]string

	Moves [][2]string

	Outcome       string
	GameString    string
	OutcomeMethod string

	IsCheck bool
	IsMate  bool

	Players []PlayerFormat

	Timer string

	Turn string
}
