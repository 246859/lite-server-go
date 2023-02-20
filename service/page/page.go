package page

import (
	"gorm.io/gorm"
	"liteserver/model/sys/sysreq"
)

// PageService
// @Date 2023-02-20 17:17:11
// @Description: 分页操作封装
type PageService struct {
}

type Where func(*gorm.DB) *gorm.DB

type Page func(model *gorm.DB, where Where, dst interface{}) *gorm.DB

// SelectPage
// @Date 2023-02-20 17:22:09
// @Param pageInfo sysreq.PageInfo
// @Return Page
// @Description: 返回一个闭包函数，调用时进行分页查询
func (PageService) SelectPage(pageInfo sysreq.PageInfo) Page {
	offset := (pageInfo.Page - 1) * pageInfo.Size
	limit := pageInfo.Size
	return func(model *gorm.DB, where Where, dst interface{}) *gorm.DB {
		db := model.Offset(offset).Limit(limit).Order(`updated_at`)
		// where执行
		db = where(db)
		// 如果是降序排序的话
		if pageInfo.Desc {
			db = db.Order("updated_at desc")
		} else {
			db = db.Order("updated_at")
		}
		return db.Find(dst)
	}
}
