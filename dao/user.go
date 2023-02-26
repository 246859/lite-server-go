package dao

import (
	"github.com/246859/lite-server-go/model"
	"github.com/246859/lite-server-go/model/request"
	"github.com/246859/lite-server-go/model/response"
	"gorm.io/gorm"
)

type UserDao struct {
	pageDao PageDao
}

// Delete
// @Date 2023-02-26 19:53:05
// @Param db *gorm.DB
// @Param id uint
// @Return error
// @Description: 软删除数据库中用户的一个数据
func (UserDao) Delete(db *gorm.DB, id uint) error {
	return db.Model(model.SystemUser{}).Delete("id = ?", id).Error
}

// UpdateBasicInfo
// @Date 2023-02-26 19:55:52
// @Param db *gorm.DB
// @Param id int
// @Param user request.UpdateUser
// @Return error
// @Description: 更新用户的基本信息
func (UserDao) UpdateBasicInfo(db *gorm.DB, id uint, user request.UpdateUser) error {
	return db.Model(model.SystemUser{}).Where("id = ?", id).UpdateColumns(user).Error
}

// GetSimpleInfo
// @Date 2023-02-26 20:00:27
// @Param db *gorm.DB
// @Param id uint
// @Return response.UserSimpleInfo
// @Return error
// @Description: 获取用户记录的简单信息
func (UserDao) GetSimpleInfo(db *gorm.DB, id uint) (response.UserSimpleInfo, error) {
	var userSimpleInfo response.UserSimpleInfo
	err := db.Model(model.SystemUser{}).Where("id = ?", id).First(&userSimpleInfo).Error
	return userSimpleInfo, err
}

// GetBasicInfo
// @Date 2023-02-26 20:00:57
// @Param db *gorm.DB
// @Param id int
// @Return response.UserBasicInfo
// @Return error
// @Description: 获取用户记录的基本信息
func (UserDao) GetBasicInfo(db *gorm.DB, id uint) (response.UserBasicInfo, error) {
	var basicInfo response.UserBasicInfo
	err := db.Model(model.SystemUser{}).Where("id = ?", id).First(&basicInfo).Error
	return basicInfo, err
}

// GetUserSimpleInfoList
// @Date 2023-02-26 20:09:41
// @Param db *gorm.DB
// @Param pageInfo request.PageInfo
// @Return []response.UserSimpleInfo
// @Return error
// @Description: 分页获取用户简单信息记录列表
func (u UserDao) GetUserSimpleInfoList(db *gorm.DB, pageInfo request.PageInfo) ([]response.UserSimpleInfo, error) {
	var simpleInfoList []response.UserSimpleInfo
	dbModel := db.Model(model.SystemUser{})
	page := u.pageDao.SelectPage(pageInfo)
	err := page(dbModel, nil, &simpleInfoList).Error
	return simpleInfoList, err
}

// GetUserBasicInfoList
// @Date 2023-02-26 20:10:05
// @Param db *gorm.DB
// @Param pageInfo request.PageInfo
// @Return []response.UserBasicInfo
// @Return error
// @Description: 分页获取用户基本信息记录列表
func (u UserDao) GetUserBasicInfoList(db *gorm.DB, pageInfo request.PageInfo) ([]response.UserBasicInfo, error) {
	var basicInfoList []response.UserBasicInfo
	dbModel := db.Model(model.SystemUser{})
	page := u.pageDao.SelectPage(pageInfo)
	err := page(dbModel, nil, &basicInfoList).Error
	return basicInfoList, err
}
