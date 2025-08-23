package model

type Data struct {
	FEN      string
	White    PlayerFormat
	Black    PlayerFormat
	Status   string
	Message  string
	LastMove string
}

type Message struct {
	Type string // CONNECTED , MOVE, RESING, START, TURN , TIMEOUT, WAIT, END
	Data Data
}
