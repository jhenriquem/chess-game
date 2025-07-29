package model

type PlayerFormat struct {
	Color    string  // W or B
	TimeLeft float64 `json:"time_left"`
	Moves    []string
	Score    int
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
