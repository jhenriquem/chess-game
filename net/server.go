package net

import (
	"chess-game/model"
	"log"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
}

func (c *Client) SendMove(move string) error {
	msg := model.Message{
		Type: "MOVE",
		Data: model.Data{
			Message: move,
		},
	}

	if err := c.Conn.WriteJSON(msg); err != nil {
		log.Println("Erro ao enviar mensagem:", err)
		return err
	}
	return nil
}

func (c *Client) ReadServer(message chan model.Message, errChan chan error) {
	for {
		var m model.Message
		if err := c.Conn.ReadJSON(&m); err != nil {
			errChan <- err
			return
		}
		message <- m
	}
}

func ConnectedServer(name string) (*Client, error) {
	u := url.URL{Scheme: "ws", Host: os.Getenv("SERVER_URL"), Path: "/"}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}

	client := &Client{
		Conn: conn,
	}

	msg := model.Message{
		Type: "CONNECTED",
		Data: model.Data{
			FEN:     "",
			Message: name,
		},
	}

	client.Conn.WriteJSON(msg)

	return client, nil
}
