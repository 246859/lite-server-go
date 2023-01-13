package sys

import "gorm.io/gorm"

// SystemUser
// @Date 2023-01-13 15:23:50
// @Description: 系统用户结构体
type SystemUser struct {
	gorm.Model
	Uuid     uint64 `gorm:"primaryKey;comment:用户ID;" json:"uuid"`
	Username string `gorm:"comment:用户登录账号;" json:"username"`
	Password string `gorm:"comment:用户登录密码" json:"password"`
	Nickname string `gorm:"comment:用户昵称;" json:"nickname"`
	Avatar   string `gorm:"comment:用户头像;" json:"avatar"`
	Phone    string `gorm:"comment:用户手机号;" json:"phone"`
	Email    string `gorm:"comment:用户邮箱;" json:"email"`
	Enable   bool   `gorm:"comment:用户是否启用;" json:"enable"`
}
