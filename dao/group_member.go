package dao

import "imgo/model"

type memberDAO struct{}

var memberInstance *memberDAO

func NewMember() *memberDAO {
	if memberInstance == nil {
		memberInstance = &memberDAO{}
	}
	return memberInstance
}

func (*memberDAO) AddMember(m *model.Group_member) error {
	if err := DB.Create(m).Error; err != nil {
		return err
	}
	return nil
}

func (*memberDAO) JoinGroup(m *model.Group_member) error {
	if err := DB.Create(m).Error; err != nil {
		return err
	}
	return nil
}

func (*memberDAO) QueryMemberByID(gid uint) ([]model.User, error) {
	var users []model.User
	sql := "SELECT u.id,u.user_name FROM `group` AS g JOIN group_member AS gm ON gm.gid = g.id JOIN user AS u ON u.id = gm.uid WHERE g.id = ?"
	if err := DB.Raw(sql, gid).Scan(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (*memberDAO) IsGroupMember(uid uint, gid uint) (bool, error) {
	var count int64
	err := DB.Model(&model.Group_member{}).Where("uid = ? AND gid = ?", uid, gid).Count(&count).Error
	if err == nil && count != 0 {
		return true, nil
	}
	return false, err
}
