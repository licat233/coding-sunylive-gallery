package logic

import (
	"net/http"
	"sunylive-gallery/errorx"
	"sunylive-gallery/global"
	"sunylive-gallery/model"
	"sunylive-gallery/respx"
	"sunylive-gallery/utils"

	"github.com/gin-gonic/gin"
)

func (l *Logic) AddUrlsPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "addurl.html", gin.H{
		"title": "Login",
	})
}

type AddUrlRequest struct {
	Urls []string `json:"urls"`
}

func (l *Logic) AddUrls(ctx *gin.Context) {
	var req = new(AddUrlRequest)
	err := ctx.BindJSON(req)
	if err != nil {
		errorx.ExternalError(ctx, err)
		return
	}
	var invalidUrls []string //不合格的url
	var validUrls []string   //合格的url
	//提取出合格的url
	for _, value := range req.Urls {
		urls, err := utils.ImgUrlPick(value)
		if err != nil {
			invalidUrls = append(invalidUrls, urls...)
			continue
		}
		validUrls = append(validUrls, urls...)
	}
	//validUrls 转化为 model.Gallery
	gallerys := []*model.Gallery{}
	for _, url := range validUrls {
		gallerys = append(gallerys, &model.Gallery{
			URL: url,
		})
	}
	//保存到数据库
	err = l.GalleryMgr.InsertsByUrl(gallerys)
	if err != nil {
		global.Log.Errorf("InsertsByUrl error: %s", err.Error())
		errorx.InternalError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, respx.SingleResp("保存成功", validUrls))
}
