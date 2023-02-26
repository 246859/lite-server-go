package system

import (
	"github.com/246859/lite-server-go/dao"
	"github.com/246859/lite-server-go/global"
	"github.com/246859/lite-server-go/model/request"
	"github.com/246859/lite-server-go/model/response"
	"github.com/246859/lite-server-go/utils"
)

type UserService struct {
	userDao dao.UserDao
}

// GetUserSimpleInfo
// @Date 2023-02-26 21:12:26
// @Param id int
// @Return response.UserSimpleInfo
// @Return error
// @Description: 查询一个用户的简单信息
func (u UserService) GetUserSimpleInfo(id uint) (response.UserSimpleInfo, error) {
	info, err := u.userDao.GetSimpleInfo(global.DB(), id)
	if err != nil {
		return response.UserSimpleInfo{}, utils.WrapSimpleError(global.I18nRawCN("user.info.fail"), err)
	}
	return info, nil
}

// GetUserBasicInfo
// @Date 2023-02-26 21:12:21
// @Param id int
// @Return response.UserBasicInfo
// @Return error
// @Description: 查询一个用户的基本信息
func (u UserService) GetUserBasicInfo(id uint) (response.UserBasicInfo, error) {
	info, err := u.userDao.GetBasicInfo(global.DB(), id)
	if err != nil {
		return response.UserBasicInfo{}, utils.WrapSimpleError(global.I18nRawCN("user.info.fail"), err)
	}
	return info, nil
}

// GetUserSimpleInfoList
// @Date 2023-02-26 21:12:17
// @Param pageInfo request.PageInfo
// @Return []response.UserSimpleInfo
// @Return error
// @Description: 查询用户简单信息列表
func (u UserService) GetUserSimpleInfoList(pageInfo request.PageInfo) ([]response.UserSimpleInfo, error) {
	list, err := u.userDao.GetUserSimpleInfoList(global.DB(), pageInfo)
	if err != nil {
		return nil, utils.WrapSimpleError(global.I18nRawCN("user.list.fail"), err)
	}
	return list, nil
}

// GetUserBasicInfoList
// @Date 2023-02-26 21:12:11
// @Param pageInfo request.PageInfo
// @Return []response.UserBasicInfo
// @Return error
// @Description: 查询用户基本信息列表
func (u UserService) GetUserBasicInfoList(pageInfo request.PageInfo) ([]response.UserBasicInfo, error) {
	list, err := u.userDao.GetUserBasicInfoList(global.DB(), pageInfo)
	if err != nil {
		return nil, utils.WrapSimpleError(global.I18nRawCN("user.list.fail"), err)
	}
	return list, nil
}

// DeleteUser
// @Date 2023-02-26 21:12:06
// @Param id int
// @Return error
// @Description: 删除一个用户
func (u UserService) DeleteUser(id uint) error {
	err := u.userDao.Delete(global.DB(), id)
	if err != nil {
		return utils.WrapSimpleError(global.I18nRawCN("user.delete.fail"), err)
	}
	return nil
}

// UpdateUserInfo
// @Date 2023-02-26 21:11:58
// @Param id int
// @Param user request.UpdateUser
// @Return error
// @Description: 更新一个用的基本信息
func (u UserService) UpdateUserInfo(id uint, user request.UpdateUser) error {
	err := u.userDao.UpdateBasicInfo(global.DB(), id, user)
	if err != nil {
		return utils.WrapSimpleError(global.I18nRawCN("user.update.fail"), err)
	}
	return nil
}
