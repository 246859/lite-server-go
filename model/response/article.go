package response

import (
	"time"
)

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

// ArticleCommentInfo
// @Date 2023-02-23 22:14:48
// @Description: 文章评论信息
type ArticleCommentInfo struct {
	UserId    uint        `json:"userId" label:"用户ID"`
	ArticleId uint        `json:"articleId" label:"文章ID"`
	ReplyList []ReplyInfo `json:"replyId" label:"回复ID"`
	Content   string      `json:"content" label:"评论内容"`
	UpdatedAt time.Time   `json:"UpdatedAt" label:"更新时间"`
}
