package logic

import (
	"sunylive-gallery/model"
)

type Logic struct {
	GalleryMgr *model.GalleryModel
}

func NewLogic() *Logic {
	return &Logic{
		GalleryMgr: model.NewGalleryModel(),
	}
}
