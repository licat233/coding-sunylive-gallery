package model

import (
	"time"
)

// Gallery 上宜居画廊
type Gallery struct {
	ID        int64     `gorm:"primaryKey;column:id" json:"-"`      // 主键
	Name      string    `gorm:"column:name" json:"name"`            // 图片名
	URL       string    `gorm:"column:url" json:"url"`              // 图片链接
	Preview   string    `gorm:"column:preview" json:"preview"`      // 图片预览链接
	Tag       string    `gorm:"column:tag" json:"tag"`              // 标签
	Alt       string    `gorm:"column:alt" json:"alt"`              // 图片alt
	CatgoryID int64     `gorm:"column:catgory_id" json:"catgoryId"` // 类别id,用于后期拓展
	IsVideo   bool      `gorm:"column:is_video" json:"isVideo"`     // 是视频？
	IsDelete  bool      `gorm:"column:is_delete" json:"isDelete"`   // 已被删除？
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
	CreateAt  time.Time `gorm:"column:create_at" json:"createAt"`   // 创建时间
	UpdateAt  time.Time `gorm:"column:update_at" json:"updateAt"`   // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *Gallery) TableName() string {
	return "gallery"
}

// GalleryColumns get sql column name.获取数据库列名
var GalleryColumns = struct {
	ID        string
	Name      string
	URL       string
	Preview   string
	Tag       string
	Alt       string
	CatgoryID string
	IsVideo   string
	IsDelete  string
	DeletedAt string
	CreateAt  string
	UpdateAt  string
}{
	ID:        "id",
	Name:      "name",
	URL:       "url",
	Preview:   "preview",
	Tag:       "tag",
	Alt:       "alt",
	CatgoryID: "catgory_id",
	IsVideo:   "is_video",
	IsDelete:  "is_delete",
	DeletedAt: "deleted_at",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
}
