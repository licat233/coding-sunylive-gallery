package server

import (
	"sunylive-gallery/server/logic"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if logic.GetUserStatus(c) == false {
			// c.JSON(http.StatusUnauthorized, respx.StateResp("请先登录"))
			c.Redirect(302, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}
