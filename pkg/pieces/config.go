package pieces

type pieces struct {
	Knight string
	King   string
	Queen  string
	Bishop string
	Tower  string
	Pawn   string
}

var BlackPieces = pieces{
	Knight: "♞",
	King:   "♚",
	Queen:  "♛",
	Bishop: "♝",
	Tower:  "♜",
	Pawn:   "♟",
}

var WhitePieces = pieces{
	Knight: "♘",
	King:   "♔",
	Queen:  "♕",
	Bishop: "♗",
	Tower:  "♖",
	Pawn:   "♙",
}
