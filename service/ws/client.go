package ws

import (
	"encoding/json"
	"imgo/pkg/util"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 6 * 60 * time.Second
	pingPeriod     = (pongWait * 9) / 1
	maxMessageSize = 512
)

// 用户实体
type Client struct {
	ID   string
	Conn *websocket.Conn
	Send chan []byte
}

type SendMsg struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Type    int    `json:"type"`
	Content string `json:"content"`
}

type ReplyMsg struct {
	From    string `json:"from"`
	Content string `json:"content"`
}

// 从ws连接中读取信息
func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		c.Conn.PongHandler()
		sendMsg := &SendMsg{}
		err := c.Conn.ReadJSON(sendMsg)
		message, _ := json.Marshal(sendMsg)

		if err != nil {
			util.LogInstance.Info("---读取客户端信息失败---")
			Manager.Unregister <- c
			c.Conn.Close()
			break
		}

		Manager.Broadcast <- message
	}

}

// 将client中信息写入websocket连接
func (c *Client) Write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
				util.LogInstance.Info("---用户读取信息失败---")
			}

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
