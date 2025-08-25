package net

import (
	"chess-game/model"
	"log"
	"net/url"

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
	serverURL := "chess-game-server-d6my.onrender.com"
	u := url.URL{Scheme: "wss", Host: serverURL, Path: "/game"}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Println(u.String())
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
