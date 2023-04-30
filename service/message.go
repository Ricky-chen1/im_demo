package service

import (
	"imgo/dao"
	"imgo/model"
)

type msgService struct{}

var msgInstance *msgService

func NewMsg() *msgService {
	if msgInstance == nil {
		msgInstance = &msgService{}
	}
	return msgInstance
}

func (*msgService) SaveMsg(time int64, from string, to string) error {
	msg := dao.NewMsg()
	newMsg := &model.Msg{
		From:       from,
		To:         to,
		ExpireTime: time,
	}

	if err := msg.CreateMsg(newMsg); err != nil {
		return err
	}
	return nil
}

func (*msgService) FindMsgsByTime(time int64, from string, to string) ([]model.Msg, error) {
	msg := dao.NewMsg()
	msgs, err := msg.FindMsgsByTime(time, from, to)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}
