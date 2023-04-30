package dto

type Common struct {
	Id uint `json:"id" form:"id"`
}

type UserRegister struct {
	UserName string `json:"user_name" form:"user_name"`
	NickName string `json:"nick_name" form:"nick_name"`
	Password string `json:"password" form:"password"`
}

type UserLogin struct {
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
}
