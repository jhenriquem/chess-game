package model

type Protoplayer struct {
	Color         string
	TimeRemaining string
	Moves         []string
	Score         int
}

type Protogame struct {
	Board [8][8]string

	Moves [][2]string

	Outcome       string
	GameString    string
	OutcomeMethod string
	MoveResult    string

	Player Protoplayer

	Timer string

	Turn string
}
