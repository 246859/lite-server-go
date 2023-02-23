package dao

import (
	"github.com/246859/lite-server-go/model/request"
	"gorm.io/gorm"
)

var PageHelper = new(PageDao)

// PageDao
// @Date 2023-02-20 17:17:11
// @Description: 分页操作封装
type PageDao struct {
}

type Where func(*gorm.DB) *gorm.DB

type Page func(model *gorm.DB, where Where, dst interface{}) *gorm.DB

// SelectPage
// @Date 2023-02-20 17:22:09
// @Param pageInfo sysreq.PageInfo
// @Return Page
// @Description: 返回一个闭包函数，调用时进行分页查询
func (PageDao) SelectPage(pageInfo request.PageInfo) Page {
	offset := (pageInfo.Page - 1) * pageInfo.Size
	limit := pageInfo.Size
	return func(model *gorm.DB, where Where, dst interface{}) *gorm.DB {
		db := model.Offset(offset).Limit(limit).Order(`updated_at`)
		if where != nil {
			// where执行
			db = where(db)
		}
		// 如果是降序排序的话
		if pageInfo.Desc != 0 {
			db = db.Order("updated_at desc")
		} else {
			db = db.Order("updated_at")
		}
		return db.Find(dst)
	}
}
