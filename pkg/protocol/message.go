package protocol

type Message struct {
	Type string `json:"type"` // "move", "info", "error"
	Data string `json:"data"` // conteúdo
}
