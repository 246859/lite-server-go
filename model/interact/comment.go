package interact

import "gorm.io/gorm"

// Comment
// @Date 2023-02-23 20:23:46
// @Description: 评论信息表
type Comment struct {
	Content string `json:"content" gorm:"comment:评论内容;type:text;"`
	gorm.Model
}
