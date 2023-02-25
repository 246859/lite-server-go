package model

import (
	"gorm.io/gorm"
)

// Article
// @Date 2023-02-18 18:03:55
// @Description: 文章数据库表
type Article struct {
	ArticleMeta
	SystemUser SystemUser `gorm:"foreignKey:UserId;constraint:onUpdate:RESTRICT,onDelete:CASCADE"`
	Class      Class      `gorm:"foreignKey:ClassId;constraint:onUpdate:RESTRICT,onDelete:CASCADE"`

	UserId  uint   `gorm:"comment:作者ID;" json:"userid"`
	ClassId uint   `gorm:"comment:分类ID;" json:"classId"`
	Title   string `gorm:"comment:文章标题;type:varchar(100);" json:"title"`
	Cover   string `gorm:"comment:文章封面;type:varchar(255);" json:"cover"`
	Label   string `gorm:"comment:文章标签;type:varchar(50);" json:"label"`
	Summary string `gorm:"comment:文章摘要;type:varchar(255);" json:"summary"`
	View    int    `gorm:"comment:浏览量;" json:"view"`
	Content string `gorm:"comment:文章内容;type:longtext;" json:"content"`
	gorm.Model
}

type ArticleMeta struct{}

func (ArticleMeta) TableComment() string {
	return "文章信息表"
}

func (ArticleMeta) TableName() string {
	return "articles"
}

// Class
// @Date 2023-02-25 19:07:01
// @Description: 文章分类表
type Class struct {
	ClassMeta
	Name string `gorm:"comment:分类名称;type:varchar(255);"`
	gorm.Model
}

type ClassMeta struct{}

func (ClassMeta) TableComment() string {
	return "分类信息表"
}

func (ClassMeta) TableName() string {
	return "classes"
}

// ArticleComment
// @Date 2023-02-23 20:21:38
// @Description: 文章评论关联表
type ArticleComment struct {
	ArticleCommentMeta
	Comment Comment `gorm:"foreignKey:CommentId;constraint:onUpdate:RESTRICT,onDelete:CASCADE"`
	Article Article `gorm:"foreignKey:ArticleId;constraint:onUpdate:RESTRICT,onDelete:CASCADE"`

	CommentId uint `json:"commentId" gorm:"comment:评论ID;primaryKey;"`
	ArticleId uint `json:"articleId" gorm:"comment:关联的文章ID;primaryKey;"`
}

type ArticleCommentMeta struct{}

func (ArticleCommentMeta) TableComment() string {
	return "文章-评论关联表"
}

func (ArticleCommentMeta) TableName() string {
	return "articles_comments"
}
