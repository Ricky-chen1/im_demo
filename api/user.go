package api

import (
	"imgo/dto"
	"imgo/pkg/errno"
	"imgo/serializer"
	"imgo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 用户注册
func UserRegister(c *gin.Context) {
	var user dto.UserRegister
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code:    errno.ParamsBindFail,
			Message: errno.CodeTag[errno.ParamsBindFail],
		})
		return
	}

	s := service.NewUser()
	newUser, err := s.Register(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code:    errno.UserCreateFail,
			Message: errno.CodeTag[errno.UserCreateFail],
		})
		return
	}

	c.JSON(http.StatusOK, serializer.Response{
		Code:    errno.Success,
		Message: errno.CodeTag[errno.Success],
		Data:    serializer.BuildUser(newUser),
	})
}

// 用户登录
func UserLogin(c *gin.Context) {
	var user dto.UserLogin
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code:    errno.ParamsBindFail,
			Message: errno.CodeTag[errno.ParamsBindFail],
		})
		return
	}

	s := service.NewUser()
	newUser, token, err := s.Login(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code:    errno.UserLoginFail,
			Message: errno.CodeTag[errno.UserLoginFail],
		})
		return
	}

	c.JSON(http.StatusOK, serializer.Response{
		Code:    errno.Success,
		Message: errno.CodeTag[errno.Success],
		Data: serializer.TokenData{
			User:  *serializer.BuildUser(newUser),
			Token: token,
		},
	})
}
