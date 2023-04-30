package api

import (
	"imgo/dto"
	"imgo/pkg/errno"
	"imgo/serializer"
	"imgo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 发送好友请求
func PushRequest(c *gin.Context) {
	var pushRequest dto.PushRequest
	if err := c.ShouldBind(&pushRequest); err != nil {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code:    errno.ParamsBindFail,
			Message: errno.CodeTag[errno.ParamsBindFail],
		})
		return
	}

	uid := c.GetUint("uid")
	err := service.NewFriend().PushRequest(uid, pushRequest.To)
	if err != nil {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code:    errno.PushRequestFail,
			Message: errno.CodeTag[errno.PushRequestFail],
		})
		return
	}

	c.JSON(http.StatusOK, serializer.Response{
		Code:    errno.Success,
		Message: errno.CodeTag[errno.Success],
	})
}

// 接受好友请求
func AcceptRequest(c *gin.Context) {
	var acceptDTO dto.AcceptRequest
	if err := c.ShouldBind(&acceptDTO); err != nil {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code:    errno.ParamsBindFail,
			Message: errno.CodeTag[errno.ParamsBindFail],
		})
		return
	}

	uid := c.GetUint("uid")
	if err := service.NewFriend().AcceptRequest(acceptDTO.ReqUid, uid); err != nil {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code:    errno.AcceptRequestFail,
			Message: errno.CodeTag[errno.AcceptRequestFail],
		})
		return
	}

	c.JSON(http.StatusOK, serializer.Response{
		Code:    errno.Success,
		Message: errno.CodeTag[errno.Success],
	})
}
