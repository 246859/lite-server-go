package model

import "gorm.io/gorm"

// SystemUser
// @Date 2023-01-13 15:23:50
// @Description: 系统用户结构体
type SystemUser struct {
	Uuid     string `gorm:"comment:用户ID;type:varchar(255);" json:"uuid" label:"用户ID"`
	Username string `gorm:"comment:用户登录账号;type:varchar(50);" json:"username" label:"用户名"`
	Password string `gorm:"comment:用户登录密码;type:varchar(255);" json:"password" label:"用户密码"`
	Nickname string `gorm:"comment:用户昵称;type:varchar(50);" json:"nickname" label:"用户昵称"`
	Avatar   string `gorm:"comment:用户头像;type:varchar(255);" json:"avatar" label:"用户头像"`
	Phone    string `gorm:"comment:用户手机号;type:varchar(50);" json:"phone" label:"用户手机号"`
	Email    string `gorm:"comment:用户邮箱;type:varchar(255);" json:"email" label:"用户邮箱"`
	Enable   bool   `gorm:"comment:用户是否启用;" json:"enable" label:"是否启用"`
	gorm.Model
}
