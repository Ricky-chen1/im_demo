package api

import (
	"imgo/dto"
	"imgo/pkg/errno"
	"imgo/serializer"
	"imgo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建群聊
func CreateGroup(c *gin.Context) {
	var group dto.CreateGroup
	if err := c.ShouldBind(&group); err != nil {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code:    errno.ParamsBindFail,
			Message: errno.CodeTag[errno.ParamsBindFail],
		})
		return
	}
	s := service.NewGroup()
	uid := c.GetUint("uid")

	if err := s.CreateGroup(group, uid); err != nil {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code:    errno.GroupCreateFail,
			Message: errno.CodeTag[errno.GroupCreateFail],
		})
		return
	}

	c.JSON(http.StatusOK, serializer.Response{
		Code:    errno.Success,
		Message: errno.CodeTag[errno.Success],
		Data:    group.Name,
	})
}

// 加入群聊
func JoinGroup(c *gin.Context) {
	var group dto.JoinGroup
	if err := c.ShouldBind(&group); err != nil {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code:    errno.ParamsBindFail,
			Message: errno.CodeTag[errno.ParamsBindFail],
		})
		return
	}

	s := service.NewGroup()
	uid := c.GetUint("uid")

	if err := s.JoinGroup(group.Gid, uid); err != nil {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code:    errno.GroupJoinFail,
			Message: errno.CodeTag[errno.GroupJoinFail],
		})
		return
	}

	c.JSON(http.StatusOK, serializer.Response{
		Code:    errno.Success,
		Message: errno.CodeTag[errno.Success],
	})
}
