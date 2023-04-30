package ws

import (
	"encoding/json"
	"imgo/pkg/util"
	"imgo/service"
	"strconv"
)

type ClientManager struct {
	Clients    map[string]*Client
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func NewManager() *ClientManager {
	return &ClientManager{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[string]*Client),
	}
}

var Manager = NewManager()

func (m *ClientManager) Start() {
	for {
		select {
		case client := <-m.Register:
			m.Clients[client.ID] = client
			replyMsg := &ReplyMsg{
				From:    "server",
				Content: "welcome " + client.ID,
			}
			//转化为byte格式进行数据传输
			msg, _ := json.Marshal(replyMsg)
			client.Send <- msg
			//再由写协程传输给客户端

		case client := <-m.Unregister:
			if _, ok := m.Clients[client.ID]; ok {
				util.LogInstance.Info(client.ID, "断开连接")

				delete(m.Clients, client.ID)
				close(client.Send)
			}

		//统一进行消息分发
		case msg := <-m.Broadcast:
			//反序列化信息
			sendMsg := &SendMsg{}
			json.Unmarshal(msg, sendMsg)

			//保存信息

			//对所有连接用户进行广播
			if sendMsg.To == "" {
				util.LogInstance.Info("send to everyone")
				for _, client := range m.Clients {
					if client.ID == sendMsg.From {
						continue
					}
					select {
					case client.Send <- msg:

					default:
						close(client.Send)
						delete(m.Clients, client.ID)
					}
				}
			}

			//单聊
			if sendMsg.Type == 1 {
				toUserID := sendMsg.To
				fromUserID := sendMsg.From
				sender := m.Clients[fromUserID]
				receiver, ok := m.Clients[toUserID]
				if ok {
					tid, _ := strconv.ParseUint(toUserID, 10, 64)
					fid, _ := strconv.ParseUint(fromUserID, 10, 64)

					//两个用户是否为好友关系
					s := service.NewFriend()
					flag := s.IsFriend(uint(fid), uint(tid))
					if !flag {
						replyMsg := &ReplyMsg{
							Content: "你与对方不是好友关系",
							From:    "server",
						}
						//将消息送回发送该消息用户
						reply_byte, _ := json.Marshal(replyMsg)
						sender.Send <- reply_byte
						return
					}
					receiver.Send <- msg
				}
				//群聊
			} else if sendMsg.Type == 2 {
				from, _ := strconv.ParseUint(sendMsg.From, 10, 64)
				to, _ := strconv.ParseUint(sendMsg.To, 10, 64)
				gs := service.NewGroup()
				users, err := gs.QueryUserByGroupID(uint(to))
				if err == nil {
					for _, user := range users {
						//不广播给发送者
						if user.ID == uint(from) {
							continue
						}

						uid := strconv.FormatUint(uint64(user.ID), 10)
						receiver, ok := m.Clients[uid]
						if !ok {
							continue
						}

						//广播信息
						if err == nil {
							receiver.Send <- msg
						}
					}
				}
			} else {
				//查找聊天信息

			}
		}
	}
}
