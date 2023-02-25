package model

import "gorm.io/gorm"

// SystemUser
// @Date 2023-01-13 15:23:50
// @Description: 系统用户结构体
type SystemUser struct {
	UserMeta
	Uuid        string `gorm:"comment:用户ID;type:varchar(255);" json:"uuid"`
	Username    string `gorm:"comment:用户登录账号;type:varchar(50);" json:"username"`
	Password    string `gorm:"comment:用户登录密码;type:varchar(255);" json:"password"`
	Nickname    string `gorm:"comment:用户昵称;type:varchar(50);" json:"nickname"`
	Description string `gorm:"comment:用户自我描述;type:varchar(100);" json:"description"`
	Avatar      string `gorm:"comment:用户头像;type:varchar(255);" json:"avatar"`
	Phone       string `gorm:"comment:用户手机号;type:varchar(50);" json:"phone"`
	Email       string `gorm:"comment:用户邮箱;type:varchar(255);" json:"email"`
	Enable      bool   `gorm:"comment:用户是否启用;" json:"enable"`
	gorm.Model
}

type UserMeta struct{}

func (UserMeta) TableComment() string {
	return "用户信息表"
}

func (UserMeta) TableName() string {
	return "users"
}
