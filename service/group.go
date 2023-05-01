package service

import (
	"imgo/dao"
	"imgo/dto"
	"imgo/model"

	"github.com/pkg/errors"
)

type groupService struct{}

var groupInstance *groupService

func NewGroup() *groupService {
	if groupInstance == nil {
		groupInstance = &groupService{}
	}
	return groupInstance
}

func (*groupService) CreateGroup(groupDTO dto.CreateGroup, ownerID uint) error {
	g := dao.NewGroup()
	gm := dao.NewMember()
	group := model.Group{
		Name:    groupDTO.Name,
		OwnerID: ownerID,
	}

	newGroup, err := g.CreateGroup(&group)
	if err != nil {
		return err
	}

	//默认群主为群成员
	member := &model.Group_member{
		Uid: ownerID,
		Gid: newGroup.ID,
	}

	if err := gm.AddMember(member); err != nil {
		return err
	}
	return nil
}

func (*groupService) QueryUserByGroupID(gid uint) ([]model.User, error) {
	gm := dao.NewMember()
	g := dao.NewGroup()
	group, err := g.FindGroup(gid)
	if err != nil {
		return nil, err
	}
	if group.ID <= 0 {
		return nil, errors.WithMessage(err, "群不存在")
	}

	users, err1 := gm.QueryMemberByID(group.ID)
	if err1 != nil {
		return nil, err
	}
	return users, nil

}

func (*groupService) IsGroupMember(uid uint, gid uint) (bool, error) {
	g := dao.NewGroup()
	gm := dao.NewMember()
	u := dao.NewUser()

	exist, err := u.IsUserExist(uid)
	if err != nil && !exist {
		return false, err
	}

	if _, err := g.FindGroup(gid); err != nil {
		return false, err
	}

	if ok, err := gm.IsGroupMember(uid, gid); !ok || err != nil {
		return false, err
	}

	return true, nil
}

func (*groupService) JoinGroup(gid uint, uid uint) error {
	gs := NewGroup()
	//暂时当作错误处理
	if ok, err := gs.IsGroupMember(uid, gid); ok {
		return err
	}

	gm := dao.NewMember()
	groupMember := &model.Group_member{
		Gid: gid,
		Uid: uid,
	}

	if err := gm.JoinGroup(groupMember); err != nil {
		return err
	}

	return nil
}
