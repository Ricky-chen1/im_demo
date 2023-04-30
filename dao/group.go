package dao

import "imgo/model"

type groupDAO struct{}

var groupInstance *groupDAO

func NewGroup() *groupDAO {
	if groupInstance == nil {
		groupInstance = &groupDAO{}
	}
	return groupInstance
}

func (*groupDAO) CreateGroup(g *model.Group) (*model.Group, error) {
	if err := DB.Create(g).Error; err != nil {
		return nil, err
	}
	return g, nil
}

//通过群id查找群
func (*groupDAO) FindGroup(gid uint) (*model.Group, error) {
	var group model.Group
	if err := DB.Model(&model.Group{}).Where("id = ?", gid).Find(&group).Error; err != nil {
		return nil, err
	}
	return &group, nil
}
