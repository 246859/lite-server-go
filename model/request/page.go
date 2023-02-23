package request

// PageInfo
// @Date 2023-02-20 17:19:33
// @Description: 分页查询结构体
type PageInfo struct {
	Page int `json:"page" form:"page" uri:"page" label:"页数" binding:"required,gte=1"`
	Size int `json:"size" form:"size" uri:"size" label:"数量" binding:"required,gte=1"`
	Desc int `json:"desc" form:"desc" uri:"desc" label:"降序" binding:""`
}
