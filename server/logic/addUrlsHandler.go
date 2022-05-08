package logic

import (
	"sunylive-gallery/errorx"
	"sunylive-gallery/global"
	"sunylive-gallery/model"

	"github.com/gin-gonic/gin"
)

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
	gallerys := []*model.Gallery{}
	for _, url := range req.Urls {
		gallerys = append(gallerys, &model.Gallery{
			URL: url,
		})
	}
	err = l.GalleryMgr.InsertsByUrl(gallerys)
	if err != nil {
		global.Log.Errorf("InsertsByUrl error: %s", err.Error())
		errorx.InternalError(ctx, err)
	}
}
