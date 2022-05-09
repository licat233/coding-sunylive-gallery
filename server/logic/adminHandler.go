package logic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (l *Logic) AdminPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin.html", gin.H{
		"title": "Login",
	})
}
