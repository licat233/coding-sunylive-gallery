package logic

import (
	"errors"
	"net/http"
	"strconv"
	"sunylive-gallery/errorx"
	"sunylive-gallery/global"
	"sunylive-gallery/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (l *Logic) Images(ctx *gin.Context) {
	//校验参数
	_size := ctx.DefaultQuery("size", "20")
	size, err := strconv.ParseInt(_size, 10, 64)
	if err != nil {
		global.Log.Errorf("size转换int64失败：%s", err.Error())
		errorx.ExternalError(ctx, errors.New("size value not number"))
		return
	}
	_current := ctx.DefaultQuery("current", "1")
	current, err := strconv.ParseInt(_current, 10, 64)
	if err != nil {
		global.Log.Errorf("current转换int64失败：%s", err.Error())
		errorx.ExternalError(ctx, errors.New("current value not number"))
		return
	}

	//查询数据
	page := model.NewPage(size, current)
	result, err := l.GalleryMgr.SelectPage(page)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.Log.Errorf("SelectPage error: %s", err.Error())
		errorx.InternalError(ctx, err)
		return
	}
	gallerys, ok := result.GetRecords().([]model.Gallery)
	if !ok {
		global.Log.Errorf("Invalid result type: %T", result)
		errorx.InternalError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gallerys)
}
