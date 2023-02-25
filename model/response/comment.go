package response

import (
	"github.com/246859/lite-server-go/model"
	"time"
)

// CommentInfo
// @Date 2023-02-23 23:02:33
// @Description: 评论展示信息
type CommentInfo struct {
	User      UserSimpleInfo `json:"user" label:"用户信息"`
	Content   string         `json:"content" label:"评论内容"`
	UpdatedAt time.Time      `json:"updatedAt" label:"最后更新时间"`
}

// ReplyInfo
// @Date 2023-02-23 23:03:03
// @Description:  回复展示信息
type ReplyInfo struct {
	model.ReplyMeta
	ID        uint           `gorm:"primaryKey;"`
	User      UserSimpleInfo `json:"user" label:"用户信息" gorm:"foreignKey:UserId"`
	UserId    uint           `json:"userId"`
	CommentId uint           `json:"commentId"`
	Content   string         `json:"content" label:"回复内容"`
	UpdatedAt time.Time      `json:"updatedAt" label:"最后更新时间"`
}
