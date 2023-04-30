package service

import (
	"imgo/dao"
	"imgo/dto"
	"imgo/model"
	"imgo/pkg/util"
)

type userService struct{}

var userInstance *userService

func NewUser() *userService {
	if userInstance == nil {
		userInstance = &userService{}
	}
	return userInstance
}

// 用户注册
func (s *userService) Register(userDTO dto.UserRegister) (*model.User, error) {
	userDAO := dao.NewUser()

	//参数校验
	if userDTO.NickName == "" {
		userDTO.NickName = userDTO.UserName
	}

	_, err := userDAO.FindUserByUserName(userDTO.UserName)
	if err != nil {
		return nil, err
	}

	newUser := model.User{
		UserName: userDTO.UserName,
		NickName: userDTO.NickName,
	}

	//密码加密
	if err := newUser.SetPassword(userDTO.Password); err != nil {
		return nil, err
	}

	user, err := userDAO.CreateUser(&newUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// 用户登录
func (s *userService) Login(userDTO dto.UserLogin) (*model.User, string, error) {
	userDAO := dao.NewUser()

	//错误处理？
	user, err := userDAO.FindUserByUserName(userDTO.UserName)
	if err != nil {
		return nil, "", err
	}

	ok, err := user.ComparePassword(userDTO.Password)
	if err != nil || !ok {
		return nil, "", err
	}

	//签发token
	token, err := util.SignToken(user.ID, user.UserName)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
