package dao

import "imgo/model"

type friendDao struct{}

var friendInstance *friendDao

func NewFriend() *friendDao {
	if friendInstance == nil {
		friendInstance = &friendDao{}
	}
	return friendInstance
}

func (*friendDao) AddFriend(friend *model.Friend) error {
	if err := DB.Create(friend).Error; err != nil {
		return err
	}
	return nil
}

func (*friendDao) IsFriend(ownerID uint, friendID uint) bool {
	var count int64
	err := DB.Model(&model.Friend{}).Where("friend_id = ? AND owner_id = ?", friendID, ownerID).
		Count(&count).Error
	if err == nil && count != 0 {
		return true
	}
	return false
}

func (*friendDao) PushRequest(req *model.Request) error {
	if err := DB.Create(req).Error; err != nil {
		return err
	}
	return nil
}

func (*friendDao) AcceptRequest(reqUid uint) error {
	if err := DB.Model(&model.Request{}).Where("req_uid = ?", reqUid).Update("is_accept", true).Error; err != nil {
		return err
	}
	return nil
}
