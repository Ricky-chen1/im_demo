package serializer

import "imgo/model"

type User struct {
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
}

type TokenData struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

func BuildUser(user *model.User) *User {
	return &User{
		UserName: user.UserName,
		NickName: user.NickName,
	}
}
