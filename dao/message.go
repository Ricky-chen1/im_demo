package dao

import (
	"github.com/pkg/errors"
	"imgo/model"
)

type msgDao struct{}

var msgInstance *msgDao

func NewMsg() *msgDao {
	if msgInstance == nil {
		msgInstance = &msgDao{}
	}
	return msgInstance
}

func (*msgDao) CreateMsg(msg *model.Msg) error {
	if err := DB.Create(msg).Error; err != nil {
		return errors.Wrap(err, "msg create fail")
	}
	return nil
}

func (*msgDao) FindMsgsByTime(expired int64, from string, to string) ([]model.Msg, error) {
	var msgs []model.Msg
	if err := DB.Model(&model.Msg{}).
		Where("expire_time > ? AND from = ? AND to = ?", expired, from, to).Find(&msgs).Error; err != nil {
		return nil, err
	}
	return msgs, nil
}
