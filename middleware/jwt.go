package middleware

import (
	"imgo/pkg/errno"
	"imgo/pkg/util"
	"imgo/serializer"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, serializer.Response{
				Code:    errno.NoToken,
				Message: errno.CodeTag[errno.NoToken],
			})
			c.Abort()
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, serializer.Response{
				Code:    errno.ParseTokenFail,
				Message: errno.CodeTag[errno.ParseTokenFail],
			})
			c.Abort()
			return
		}
		if claims.ExpiresAt < time.Now().Unix() {
			c.JSON(http.StatusUnauthorized, serializer.Response{
				Code:    errno.TokenExpired,
				Message: errno.CodeTag[errno.TokenExpired],
			})
			c.Abort()
			return
		}
		c.Set("uid", claims.ID)
		c.Next()
	}
}
