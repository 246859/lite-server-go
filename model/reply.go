package model

import (
	"gorm.io/gorm"
)

// Reply
// @Date 2023-02-23 22:30:16
// @Description: 回复信息表
type Reply struct {
	User    SystemUser `gorm:"foreignKey:UserId;constraint:onUpdate:RESTRICT,onDelete:CASCADE"`
	UserId  uint       `json:"userId" gorm:"comment:用户ID;"`
	Content string     `json:"content" gorm:"comment:回复内容;type:text;"`
	gorm.Model
}
