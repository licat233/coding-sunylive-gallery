package logic

import (
	"errors"
	"net/http"
	"sunylive-gallery/errorx"
	"sunylive-gallery/global"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (l *Logic) LoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}

func (l *Logic) Login(ctx *gin.Context) {
	usn := ctx.PostForm("username")
	pwd := ctx.PostForm("password")
	if string(usn) != global.Config.Passport || string(pwd) != global.Config.Password {
		global.Log.Errorf("Login error, password:%s | password: %s", usn, pwd)
		errorx.ExternalError(ctx, errors.New("username or password is incorrect"))
		return
	}
	sessions.Default(ctx).Set("isLogin", true)
	ctx.Redirect(http.StatusMovedPermanently, "/")
	return
}

func (l *Logic) LoginVerify(ctx *gin.Context) bool {
	if v, ok := sessions.Default(ctx).Get("isLogin").(bool); ok {
		if v {
			return true
		}
	}
	return false
}
