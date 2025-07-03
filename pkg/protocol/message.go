package protocol

type Player struct {
	Color         string
	TimeRemaining string
	Moves         []string
	Score         int
}

type pieces struct {
	Piece    string
	Location string
	Color    string
}

type Game struct {
	Board [8][8]string

	Moves [][2]string

	Players [2]Player

	Timer string

	Turn string
}

type Message struct {
	Info string `json:"info"`
	Game Game   `json:"game,omitempty"`
}
