package model

type Data struct {
	FEN      string
	Player   PlayerFormat
	Oponnent PlayerFormat
	Message  string
}

type Message struct {
	Type string // CONNECTED , MOVE, RESING, START, TURN , TIMEOUT, WAIT
	Data Data
}
