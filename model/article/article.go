package article

import (
	"gorm.io/gorm"
	"liteserver/model/sys"
)

// Article
// @Date 2023-02-18 18:03:55
// @Description: 文章数据库表
type Article struct {
	gorm.Model
	SystemUser sys.SystemUser `gorm:"foreignKey:UserId"`
	UserId     uint           `gorm:"comment:作者ID;" json:"userid" label:"作者Id" validate:"required"`
	Title      string         `gorm:"comment:文章标题;type:varchar(100);" json:"title" label:"文章标题" validate:"required"`
	Cover      string         `gorm:"comment:文章封面;type:varchar(50);" json:"cover" label:"文章封面"`
	Author     string         `gorm:"comment:文章作者;type:varchar(50);" json:"author" label:"文章作者" validate:"required"`
	Label      string         `gorm:"comment:文章标签;type:varchar(50);" json:"label" label:"文章标签" validate:"required"`
	Summary    string         `gorm:"comment:文章摘要;type:varchar(255);" json:"summary" label:"文章摘要" validate:"required"`
	View       int            `gorm:"comment:浏览量;" json:"view" label:"文章浏览量"`
	Content    string         `gorm:"comment:文章内容;type:longtext;" json:"content" label:"文章内容" validate:"required"`
}

// HeadInfo
// @Date 2023-02-18 18:04:04
// @Description: 文章简单信息
type HeadInfo struct {
	Title   string `json:"title" label:"文章标题" validate:"required"`
	Cover   string `json:"cover" label:"文章封面"`
	Author  string `json:"author" label:"文章作者" validate:"required"`
	Label   string `json:"label" label:"文章标签" validate:"required"`
	Summary string `json:"summary" label:"文章摘要" validate:"required"`
	View    int    `json:"view" label:"文章浏览量"`
}
