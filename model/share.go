package model

import "gorm.io/gorm"

// Share
// @Date 2023-02-23 20:58:37
// @Description: 动态信息表
type Share struct {
	User    SystemUser `gorm:"foreignKey:UserId;constraint:onUpdate:RESTRICT,onDelete:CASCADE"`
	UserId  uint       `json:"userId" gorm:"comment:用户ID;"`
	Content string     `json:"content" gorm:"comment:动态内容;type:text;"`
	gorm.Model
}
