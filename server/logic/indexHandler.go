package logic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (l *Logic) IndexPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Home",
	})
}
