package net

import (
	"chess-game/model"
	"encoding/json"
	"log"
	"net"
)

type Client struct {
	Conn net.Conn
	Enc  *json.Encoder
	Dec  *json.Decoder
}

func (c *Client) SendMove(move string) error {
	msg := model.Message{
		Type: "MOVE",
		Data: model.Data{
			Message: move,
		},
	}

	if err := c.Enc.Encode(msg); err != nil {
		log.Println("Erro ao enviar mensagem:", err)
		return err
	}
	return nil
}

func (c *Client) ReadServer(message chan model.Message, errChan chan error) {
	for {
		var m model.Message
		if err := c.Dec.Decode(&m); err != nil {
			errChan <- err
			return
		}
		message <- m
	}
}

func ConnectedServer(name string) (*Client, error) {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		return nil, err
	}

	client := &Client{
		Conn: conn,
		Enc:  json.NewEncoder(conn),
		Dec:  json.NewDecoder(conn),
	}

	msg := model.Message{
		Type: "CONNECTED",
		Data: model.Data{
			FEN:     "",
			Message: name,
		},
	}
	client.Enc.Encode(msg)

	return client, nil
}
