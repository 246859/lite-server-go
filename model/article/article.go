package article

import (
	"gorm.io/gorm"
	"liteserver/model/sys"
	"time"
)

// Article
// @Date 2023-02-18 18:03:55
// @Description: 文章数据库表
type Article struct {
	SystemUser sys.SystemUser `gorm:"foreignKey:UserId"`
	UserId     uint           `gorm:"comment:作者ID;" json:"userid" label:"作者Id"`
	Title      string         `gorm:"comment:文章标题;type:varchar(100);" json:"title" label:"文章标题" binding:"required"`
	Cover      string         `gorm:"comment:文章封面;type:varchar(255);" json:"cover" label:"文章封面"`
	Label      string         `gorm:"comment:文章标签;type:varchar(50);" json:"label" label:"文章标签" binding:"required"`
	Summary    string         `gorm:"comment:文章摘要;type:varchar(255);" json:"summary" label:"文章摘要" binding:"required"`
	View       int            `gorm:"comment:浏览量;" json:"view" label:"文章浏览量"`
	Content    string         `gorm:"comment:文章内容;type:longtext;" json:"content" label:"文章内容" binding:"required"`
	gorm.Model
}

// HeadInfo
// @Date 2023-02-18 18:04:04
// @Description: 文章简单信息
type HeadInfo struct {
	Id        uint      `json:"id" label:"文章ID"`
	Title     string    `json:"title" label:"文章标题"`
	Cover     string    `json:"cover" label:"文章封面"`
	Author    string    `json:"author" label:"文章作者"`
	Label     string    `json:"label" label:"文章标签"`
	Summary   string    `json:"summary" label:"文章摘要"`
	View      int       `json:"view" label:"文章浏览量"`
	Like      int       `json:"like" label:"点赞数"`
	Comment   int       `json:"comment" label:"评论数"`
	UpdatedAt time.Time `json:"updatedAt" label:"更新时间"`
}

// ArticleDetails
// @Date 2023-02-21 19:55:24
// @Description: 文章细节信息
type ArticleDetails struct {
	Id        uint      `json:"id" label:"文章ID"`
	UserId    uint      `json:"userId" label:"作者ID"`
	Author    string    `json:"author" label:"文章作者"`
	Title     string    `json:"title" label:"文章标题"`
	Cover     string    `json:"cover" label:"文章封面"`
	Label     string    `json:"label" label:"文章标签"`
	Summary   string    `json:"summary" label:"文章摘要"`
	View      int       `json:"view" label:"文章浏览量"`
	Like      int       `json:"like" label:"点赞数"`
	Comment   int       `json:"comment" label:"评论数"`
	Content   string    `json:"content" label:"文章内容"`
	UpdatedAt time.Time `json:"updatedAt" label:"更新时间"`
}
