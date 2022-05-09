package logic

import (
	"errors"
	"net/http"
	"sunylive-gallery/errorx"
	"sunylive-gallery/global"
	"sunylive-gallery/respx"

	"github.com/gin-gonic/gin"
)

func (l *Logic) LoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l *Logic) Login(ctx *gin.Context) {
	var req = new(LoginRequest)
	err := ctx.BindJSON(req)
	if err != nil {
		errorx.ExternalError(ctx, err)
		return
	}
	if req.Username != global.Config.Username || req.Password != global.Config.Password {
		global.Log.Errorf("Login error, username:%s | password: %s", req.Username, req.Password)
		errorx.ExternalError(ctx, errors.New("username or password is incorrect"))
		return
	}
	SetUserStatus(ctx, true)

	ctx.JSON(http.StatusOK, respx.StateResp("登录成功"))
	ctx.Abort()
	return
}
