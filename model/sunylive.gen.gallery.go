package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _GalleryMgr struct {
	*_BaseMgr
}

// GalleryMgr open func
func GalleryMgr(db *gorm.DB) *_GalleryMgr {
	if db == nil {
		panic(fmt.Errorf("GalleryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GalleryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gallery"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GalleryMgr) GetTableName() string {
	return "gallery"
}

// Reset 重置gorm会话
func (obj *_GalleryMgr) Reset() *_GalleryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_GalleryMgr) Get() (result Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GalleryMgr) Gets() (results []*Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GalleryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gallery{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_GalleryMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 图片名
func (obj *_GalleryMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithURL url获取 图片链接
func (obj *_GalleryMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// WithPreview preview获取 图片预览链接
func (obj *_GalleryMgr) WithPreview(preview string) Option {
	return optionFunc(func(o *options) { o.query["preview"] = preview })
}

// WithTag tag获取 标签
func (obj *_GalleryMgr) WithTag(tag string) Option {
	return optionFunc(func(o *options) { o.query["tag"] = tag })
}

// WithAlt alt获取 图片alt
func (obj *_GalleryMgr) WithAlt(alt string) Option {
	return optionFunc(func(o *options) { o.query["alt"] = alt })
}

// WithCatgoryID catgory_id获取 类别id,用于后期拓展
func (obj *_GalleryMgr) WithCatgoryID(catgoryID int64) Option {
	return optionFunc(func(o *options) { o.query["catgory_id"] = catgoryID })
}

// WithIsVideo is_video获取 是视频？
func (obj *_GalleryMgr) WithIsVideo(isVideo bool) Option {
	return optionFunc(func(o *options) { o.query["is_video"] = isVideo })
}

// WithIsDelete is_delete获取 已被删除？
func (obj *_GalleryMgr) WithIsDelete(isDelete bool) Option {
	return optionFunc(func(o *options) { o.query["is_delete"] = isDelete })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_GalleryMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateAt create_at获取 创建时间
func (obj *_GalleryMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取 更新时间
func (obj *_GalleryMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// GetByOption 功能选项模式获取
func (obj *_GalleryMgr) GetByOption(opts ...Option) (result Gallery, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GalleryMgr) GetByOptions(opts ...Option) (results []*Gallery, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_GalleryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gallery, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键
func (obj *_GalleryMgr) GetFromID(id int64) (result Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_GalleryMgr) GetBatchFromID(ids []int64) (results []*Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 图片名
func (obj *_GalleryMgr) GetFromName(name string) (results []*Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 图片名
func (obj *_GalleryMgr) GetBatchFromName(names []string) (results []*Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromURL 通过url获取内容 图片链接
func (obj *_GalleryMgr) GetFromURL(url string) (results []*Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`url` = ?", url).Find(&results).Error

	return
}

// GetBatchFromURL 批量查找 图片链接
func (obj *_GalleryMgr) GetBatchFromURL(urls []string) (results []*Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`url` IN (?)", urls).Find(&results).Error

	return
}

// GetFromPreview 通过preview获取内容 图片预览链接
func (obj *_GalleryMgr) GetFromPreview(preview string) (results []*Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`preview` = ?", preview).Find(&results).Error

	return
}

// GetBatchFromPreview 批量查找 图片预览链接
func (obj *_GalleryMgr) GetBatchFromPreview(previews []string) (results []*Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`preview` IN (?)", previews).Find(&results).Error

	return
}

// GetFromTag 通过tag获取内容 标签
func (obj *_GalleryMgr) GetFromTag(tag string) (results []*Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`tag` = ?", tag).Find(&results).Error

	return
}

// GetBatchFromTag 批量查找 标签
func (obj *_GalleryMgr) GetBatchFromTag(tags []string) (results []*Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`tag` IN (?)", tags).Find(&results).Error

	return
}

// GetFromAlt 通过alt获取内容 图片alt
func (obj *_GalleryMgr) GetFromAlt(alt string) (results []*Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`alt` = ?", alt).Find(&results).Error

	return
}

// GetBatchFromAlt 批量查找 图片alt
func (obj *_GalleryMgr) GetBatchFromAlt(alts []string) (results []*Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`alt` IN (?)", alts).Find(&results).Error

	return
}

// GetFromCatgoryID 通过catgory_id获取内容 类别id,用于后期拓展
func (obj *_GalleryMgr) GetFromCatgoryID(catgoryID int64) (results []*Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`catgory_id` = ?", catgoryID).Find(&results).Error

	return
}

// GetBatchFromCatgoryID 批量查找 类别id,用于后期拓展
func (obj *_GalleryMgr) GetBatchFromCatgoryID(catgoryIDs []int64) (results []*Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`catgory_id` IN (?)", catgoryIDs).Find(&results).Error

	return
}

// GetFromIsVideo 通过is_video获取内容 是视频？
func (obj *_GalleryMgr) GetFromIsVideo(isVideo bool) (results []*Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`is_video` = ?", isVideo).Find(&results).Error

	return
}

// GetBatchFromIsVideo 批量查找 是视频？
func (obj *_GalleryMgr) GetBatchFromIsVideo(isVideos []bool) (results []*Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`is_video` IN (?)", isVideos).Find(&results).Error

	return
}

// GetFromIsDelete 通过is_delete获取内容 已被删除？
func (obj *_GalleryMgr) GetFromIsDelete(isDelete bool) (results []*Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`is_delete` = ?", isDelete).Find(&results).Error

	return
}

// GetBatchFromIsDelete 批量查找 已被删除？
func (obj *_GalleryMgr) GetBatchFromIsDelete(isDeletes []bool) (results []*Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`is_delete` IN (?)", isDeletes).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_GalleryMgr) GetFromDeletedAt(deletedAt time.Time) (results []*Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_GalleryMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_GalleryMgr) GetFromCreateAt(createAt time.Time) (results []*Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`create_at` = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量查找 创建时间
func (obj *_GalleryMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`create_at` IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容 更新时间
func (obj *_GalleryMgr) GetFromUpdateAt(updateAt time.Time) (results []*Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`update_at` = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量查找 更新时间
func (obj *_GalleryMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`update_at` IN (?)", updateAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_GalleryMgr) FetchByPrimaryKey(id int64) (result Gallery, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gallery{}).Where("`id` = ?", id).First(&result).Error

	return
}
