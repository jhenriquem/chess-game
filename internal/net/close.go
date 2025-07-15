package net

import "github.com/gorilla/websocket"

func CloseConnection(conn *websocket.Conn) {
	conn.WriteMessage(websocket.CloseMessage, []byte{})
	conn.Close()
}
