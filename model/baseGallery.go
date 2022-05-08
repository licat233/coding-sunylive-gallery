package model

import (
	"errors"
	"sunylive-gallery/global"

	"gorm.io/gorm"
)

type GalleryModel struct {
	*_GalleryMgr
}

func NewGalleryModel() *GalleryModel {
	return &GalleryModel{
		_GalleryMgr: GalleryMgr(global.DB),
	}
}

func (m *GalleryModel) baseModel() *gorm.DB {
	return m.DB.WithContext(m.ctx).Model(Gallery{})
}

func (m *GalleryModel) Insert(gallery *Gallery) error {
	return m.baseModel().Omit("id", "deleted_at", "create_at", "update_at").Create(gallery).Error
}

//查询是否存在
func (m *GalleryModel) IsExist(url string) (bool, error) {
	err := m.baseModel().Where("url = ?", url).First(Gallery{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		global.Log.Errorf("gorm查询数据出错 IsExist error: %s", err.Error())
		return false, err
	}
	return true, nil
}

//插入，通过url查询是否存在，如果存在，则不再插入
func (m *GalleryModel) InsertByUrl(gallery *Gallery) error {
	exist, err := m.IsExist(gallery.URL)
	if err != nil {
		return err
	}
	if exist {
		//存在，不存储
		return nil
	}
	return m.Insert(gallery)
}

//批量插入，通过url查询是否存在，如果存在，则不再插入
func (m *GalleryModel) InsertsByUrl(gallery []*Gallery) error {
	for _, g := range gallery {
		err := m.InsertByUrl(g)
		if err != nil {
			return err
		}
	}
	return nil
}
