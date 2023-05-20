package main

import "github.com/gorilla/websocket"

// clientはチャットを行っている1人のユーザを表す
type client struct {
	// socketはこのクライアントのためのWebSocket
	socket *websocket.Conn
	// sendはこのメッセージが送られるチャネル。キューとして蓄積され、WebSocketを通じてユーザのブラウザに送られるのを待機する
	send chan []byte
	// roomはこのクライアントが参加しているチャットルーム
	room *room
}

func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

// c.sendの内容をwebsocketに書き込む
func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
