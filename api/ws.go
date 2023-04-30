package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"imgo/pkg/util"
	"imgo/service/ws"
	"net/http"
)

var upgrater = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// websocket连接
func WSServer(c *gin.Context) {
	uid := c.Query("uid")
	conn, err := upgrater.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		util.LogInstance.Info(err)
		return
	}

	client := &ws.Client{ID: uid, Conn: conn, Send: make(chan []byte, 256)}
	util.LogInstance.Info("---有客户端连接---")

	//客户端连接
	ws.Manager.Register <- client
	//开启两个协程进行读写
	go client.Write()
	go client.Read()
}
