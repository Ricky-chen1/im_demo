package service

import (
	"imgo/dao"
	"imgo/model"
	"time"

	"github.com/pkg/errors"
)

type friendService struct{}

var friendInstance *friendService

func NewFriend() *friendService {
	if friendInstance == nil {
		friendInstance = &friendService{}
	}
	return friendInstance
}

// 发送好友请求
func (s *friendService) PushRequest(from uint, to uint) error {
	f := dao.NewFriend()
	u := dao.NewUser()
	if from == to {
		return errors.New("不能向自己发送好友请求")
	}

	if exist, err := u.IsUserExist(to); !exist || err != nil {
		return err
	}

	req := &model.Request{
		To:     to,
		ReqUid: from,
	}

	if err := f.PushRequest(req); err != nil {
		return err
	}
	return nil
}

// 接受好友请求
func (s *friendService) AcceptRequest(reqUid uint, uid uint) error {
	f := dao.NewFriend()

	if err := f.AcceptRequest(reqUid); err != nil {
		return err
	}

	//创建好友关系(双向关系)
	if err := s.AddFriend(reqUid, uid); err != nil {
		return err
	}

	if err := s.AddFriend(uid, reqUid); err != nil {
		return err
	}
	return nil
}

// 添加好友关系
func (s *friendService) AddFriend(ownerID uint, friendID uint) error {
	f := dao.NewFriend()

	if f.IsFriend(ownerID, friendID) {
		return errors.New("已经为好友关系")
	}

	friend := model.Friend{
		OwnerID:    ownerID,
		FriendID:   friendID,
		CreateTime: time.Now().Unix(),
	}

	if err := f.AddFriend(&friend); err != nil {
		return err
	}

	return nil
}

// 是否为好友
func (s *friendService) IsFriend(from uint, to uint) bool {
	f := dao.NewFriend()

	return f.IsFriend(from, to)
}
