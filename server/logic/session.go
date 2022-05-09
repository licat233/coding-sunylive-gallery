package logic

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetUserStatus(c *gin.Context, status bool) {
	sess := sessions.Default(c)
	sess.Set("isLogin", status)
	//一定要记得保存
	sess.Save()
}

func GetUserStatus(c *gin.Context) bool {
	value := sessions.Default(c).Get("isLogin")
	if v, ok := value.(bool); ok {
		return v
	}
	return false
}
