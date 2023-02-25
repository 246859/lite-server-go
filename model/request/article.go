package request

type PostArticle struct {
	Title   string `json:"title" label:"文章标题" binding:"required"`
	Cover   string `json:"cover" label:"文章封面"`
	Label   string `json:"label" label:"文章标签" binding:"required"`
	Summary string `json:"summary" label:"文章摘要" binding:"required"`
	Content string `json:"content" label:"文章内容" binding:"required"`
}
