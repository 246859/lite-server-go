package dao

import (
	"github.com/246859/lite-server-go/model"
	"gorm.io/gorm"
)

type ClassDao struct {
}

// GetOne
// @Date 2023-02-27 22:00:52
// @Param db *gorm.DB
// @Param names string
// @Return model.Class
// @Return error
// @Description: 根据名称查分类
func (ClassDao) GetOne(db *gorm.DB, names string) (model.Class, error) {
	var class model.Class
	err := db.Model(model.Class{}).Where("name = ?", names).First(&class).Error
	return class, err
}
