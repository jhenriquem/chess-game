package model

type Data struct {
	FEN     string
	Message string
}

type Message struct {
	Type string // CONNECTED , MOVE, RESING, START, TURN , TIMEOUT
	Data Data
}
