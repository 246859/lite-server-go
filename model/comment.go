package model

import (
	"gorm.io/gorm"
)

// Comment
// @Date 2023-02-23 20:23:46
// @Description: 评论信息表
type Comment struct {
	User    SystemUser `gorm:"foreignKey:UserId;constraint:onUpdate:RESTRICT,onDelete:CASCADE"`
	UserId  uint       `json:"userId" gorm:"comment:用户ID;"`
	Content string     `json:"content" gorm:"comment:评论内容;type:text;"`
	gorm.Model
}

// CommentReply
// @Date 2023-02-23 22:30:59
// @Description: 评论回复关联表
type CommentReply struct {
	Comment   Comment `gorm:"foreignKey:CommentId;constraint:onUpdate:RESTRICT,onDelete:CASCADE"`
	Reply     Reply   `gorm:"foreignKey:ReplyId;constraint:onUpdate:RESTRICT,onDelete:CASCADE"`
	CommentId uint    `json:"commentId" gorm:"comment:评论ID;primaryKey;"`
	ReplyId   uint    `json:"replyId" gorm:"comment:回复ID;primaryKey;"`
}
