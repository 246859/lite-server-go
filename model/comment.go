package model

import (
	"gorm.io/gorm"
)

// Comment
// @Date 2023-02-23 20:23:46
// @Description: 评论信息表
type Comment struct {
	CommentTN
	User SystemUser `gorm:"foreignKey:UserId;constraint:onUpdate:RESTRICT,onDelete:CASCADE"`

	UserId  uint   `json:"userId" gorm:"comment:用户ID;"`
	Content string `json:"content" gorm:"comment:评论内容;type:text;"`
	gorm.Model
}

type CommentTN struct{}

func (CommentTN) TableName() string {
	return "comments"
}
