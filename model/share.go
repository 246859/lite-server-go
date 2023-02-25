package model

import "gorm.io/gorm"

// Share
// @Date 2023-02-23 20:58:37
// @Description: 动态信息表
type Share struct {
	ShareMeta
	User SystemUser `gorm:"foreignKey:UserId;constraint:onUpdate:RESTRICT,onDelete:CASCADE"`

	UserId  uint   `json:"userId" gorm:"comment:用户ID;"`
	Content string `json:"content" gorm:"comment:动态内容;type:text;"`
	gorm.Model
}

type ShareMeta struct{}

func (ShareMeta) TableComment() string {
	return "动态信息表"
}

func (ShareMeta) TableName() string {
	return "shares"
}

// ShareComment
// @Date 2023-02-24 23:44:43
// @Description: 动态评论关联表
type ShareComment struct {
	ShareCommentMeta
	Comment Comment `gorm:"foreignKey:CommentId;constraint:onUpdate:RESTRICT,onDelete:CASCADE"`
	Share   Share   `gorm:"foreignKey:ShareId;constraint:onUpdate:RESTRICT,onDelete:CASCADE"`

	CommentId uint `json:"commentId" gorm:"comment:评论ID;primaryKey;"`
	ShareId   uint `json:"shareId" gorm:"comment:动态ID;primaryKey;"`
}

type ShareCommentMeta struct{}

func (ShareCommentMeta) TableComment() string {
	return "动态-评论关联表"
}

func (ShareCommentMeta) TableName() string {
	return "shares_comments"
}
