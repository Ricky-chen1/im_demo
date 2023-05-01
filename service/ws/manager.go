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

			msgS := service.NewMsg()
			//保存信息(错误处理?)
			msgS.SaveMsg(sendMsg.From, sendMsg.To)
			//设置消息过期时间?

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
				//查找聊天信息(此时内容传递时间)
				msgS := service.NewMsg()
				time, _ := strconv.ParseInt(sendMsg.Content, 10, 64)
				timeString := util.ConvUnixToTime(time)

				//按照发送时间查询
				msgs, err := msgS.FindMsgsByTime(timeString, sendMsg.From, sendMsg.To)
				sender := m.Clients[sendMsg.From]
				if err == nil {
					for _, item := range msgs {
						msgByte := []byte(item.Content)

						//聊天内容实时返回
						sender.Send <- msgByte
					}
					//传输完毕后关闭通道
					close(sender.Send)
					delete(m.Clients, sender.ID)
				}
			}
		}
	}
}
