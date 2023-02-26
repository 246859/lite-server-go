package response

import (
	"github.com/246859/lite-server-go/model"
	"time"
)

// UserSimpleInfo
// @Date 2023-02-23 23:00:14
// @Description: 用户简单信息
type UserSimpleInfo struct {
	model.UserMeta `json:"-"`
	Id             uint   `json:"id" label:"用户id"`
	Avatar         string `json:"avatar" label:"头像"`
	Email          string `json:"email" label:"用户邮箱"`
	Nickname       string `json:"nickname" label:"昵称"`
}

// UserBasicInfo
// @Date 2023-02-25 19:20:03
// @Description: 用户基本信息结构体
type UserBasicInfo struct {
	model.UserMeta `json:"-"`
	Id             uint      `json:"id" label:"用户id"`
	Uuid           string    `json:"uuid" label:"用户唯一ID"`
	Avatar         string    `json:"avatar" label:"头像"`
	Email          string    `json:"email" label:"用户邮箱"`
	Nickname       string    `json:"nickname" label:"昵称"`
	Description    string    `json:"description" label:"自我描述"`
	Password       string    `json:"password" label:"密码"`
	CreatedAt      time.Time `json:"createdAt" label:"注册时间"`
}
