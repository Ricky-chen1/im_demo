package dao

import (
	"github.com/pkg/errors"
	"imgo/model"
)

type userDAO struct{}

var userInstance *userDAO

func NewUser() *userDAO {
	if userInstance == nil {
		userInstance = &userDAO{}
	}
	return userInstance
}

// 根据用户名查找用户
func (*userDAO) FindUserByUserName(username string) (*model.User, error) {
	var user model.User
	if err := DB.Model(&model.User{}).Where("user_name = ?", username).Find(&user).Error; err != nil {
		return nil, errors.Wrap(err, "query err")
	}
	return &user, nil
}

// 创建用户
func (*userDAO) CreateUser(user *model.User) (*model.User, error) {
	if err := DB.Create(user).Error; err != nil {
		return nil, errors.Wrap(err, "create err")
	}
	return user, nil
}

func (*userDAO) IsUserExist(uid uint) (bool, error) {
	var count int64
	err := DB.Model(&model.User{}).Where("id = ?", uid).Count(&count).Error
	if err == nil && count != 0 {
		return true, err
	}
	return false, nil
}
