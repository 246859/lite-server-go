package model

import (
	"gorm.io/gorm"
)

// Like
// @Date 2023-02-23 20:25:46
// @Description: 点赞信息表
type Like struct {
	User   SystemUser `gorm:"foreignKey:UserId;constraint:onUpdate:RESTRICT,onDelete:CASCADE"`
	UserId uint       `json:"userId" gorm:"comment:用户ID;"`
	Like   int        `json:"like" gorm:"comment:点赞数量;type:bigint;"`
	gorm.Model
}
