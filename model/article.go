package model

import (
	"gorm.io/gorm"
)

// Article
// @Date 2023-02-18 18:03:55
// @Description: 文章数据库表
type Article struct {
	ArticleTN
	SystemUser SystemUser `gorm:"foreignKey:UserId;constraint:onUpdate:RESTRICT,onDelete:CASCADE"`

	UserId  uint   `gorm:"comment:作者ID;" json:"userid" label:"作者Id"`
	Title   string `gorm:"comment:文章标题;type:varchar(100);" json:"title" label:"文章标题" binding:"required"`
	Cover   string `gorm:"comment:文章封面;type:varchar(255);" json:"cover" label:"文章封面"`
	Label   string `gorm:"comment:文章标签;type:varchar(50);" json:"label" label:"文章标签" binding:"required"`
	Summary string `gorm:"comment:文章摘要;type:varchar(255);" json:"summary" label:"文章摘要" binding:"required"`
	View    int    `gorm:"comment:浏览量;" json:"view" label:"文章浏览量"`
	Content string `gorm:"comment:文章内容;type:longtext;" json:"content" label:"文章内容" binding:"required"`
	gorm.Model
}

type ArticleTN struct{}

func (ArticleTN) TableName() string {
	return "articles"
}

// ArticleComment
// @Date 2023-02-23 20:21:38
// @Description: 文章评论关联表
type ArticleComment struct {
	Comment Comment `gorm:"foreignKey:CommentId;constraint:onUpdate:RESTRICT,onDelete:CASCADE"`
	Article Article `gorm:"foreignKey:ArticleId;constraint:onUpdate:RESTRICT,onDelete:CASCADE"`
	ArticleCommentTN

	CommentId uint `json:"commentId" gorm:"comment:评论ID;primaryKey;"`
	ArticleId uint `json:"articleId" gorm:"comment:关联的文章ID;primaryKey;"`
}

type ArticleCommentTN struct{}

func (ArticleCommentTN) TableName() string {
	return "articles_comments"
}
