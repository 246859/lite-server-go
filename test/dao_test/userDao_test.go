package dao_test

import (
	"fmt"
	"github.com/246859/lite-server-go/dao"
	"github.com/246859/lite-server-go/model/request"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
)

var db *gorm.DB

var userDao = dao.UserDao{}

func init() {
	open, err := gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?%s",
		"root", "wyh246859", "192.168.48.134", "3306", "liteserver", "charset=utf8mb4&parseTime=True&loc=Local")))
	if err != nil {
		panic(err)
	} else {
		db = open
	}
	db = open.Debug()
}

// TestGetUserSimpleInfo
// @Date 2023-02-26 20:22:39
// @Param t *testing.T
// @Description: 测试 userDao.GetSimpleInfo
func TestGetUserSimpleInfo(t *testing.T) {
	info, err := userDao.GetSimpleInfo(db, 1)
	if err != nil {
		t.Error("GetSimpleInfo测试失败", err)
	} else {
		log.Printf("%+v", info)
	}
}

// TestGetUserBasicInfo
// @Date 2023-02-26 20:26:52
// @Param t *testing.T
// @Description: 测试 userDao.GetBasicInfo
func TestGetUserBasicInfo(t *testing.T) {
	info, err := userDao.GetBasicInfo(db, 1)
	if err != nil {
		t.Error("GetBasicInfo测试失败", err)
	} else {
		log.Printf("%+v", info)
	}
}

// TestGetUserSimpleInfoList
// @Date 2023-02-26 20:49:56
// @Param t *testing.T
// @Description: 测试 userDao.GetUserSimpleInfoList
func TestGetUserSimpleInfoList(t *testing.T) {
	list, err := userDao.GetUserSimpleInfoList(db, request.PageInfo{
		Page: 1,
		Size: 2,
		Desc: 0,
	})
	if err != nil {
		t.Error(err)
	} else {
		log.Printf("%+v\n%d", list, len(list))
	}
}

// TestGetUserBasicInfoList
// @Date 2023-02-26 20:32:10
// @Param t *testing.T
// @Description: 测试 userDao.GetUserBasicInfoList
func TestGetUserBasicInfoList(t *testing.T) {
	list, err := userDao.GetUserBasicInfoList(db, request.PageInfo{
		Page: 1,
		Size: 2,
		Desc: 0,
	})
	if err != nil {
		t.Error(err)
	} else {
		log.Printf("%+v\n%d", list, len(list))
	}
}

// TestDeleteUser
// @Date 2023-02-26 20:35:22
// @Param t *testing.T
// @Description: 测试 userDao.Delete
func TestDeleteUser(t *testing.T) {
	err := userDao.Delete(db, 1)
	if err != nil {
		t.Error(err)
	}
}

// TestUpdateUser
// @Date 2023-02-26 20:40:36
// @Param t *testing.T
// @Description: 测试 userDao.UpdateBasicInfo
func TestUpdateUser(t *testing.T) {
	err := userDao.UpdateBasicInfo(db, 2, request.UpdateUser{
		Avatar:      "123456789",
		Nickname:    "giaogiao",
		Password:    "789",
		Description: "666",
	})
	if err != nil {
		t.Error(err)
	}
}
