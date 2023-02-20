package sysreq

// PageInfo
// @Date 2023-02-20 17:19:33
// @Description: 分页查询结构体
type PageInfo struct {
	Page int  `json:"page" label:"页数" validate:"required,gte=1"`
	Size int  `json:"size" label:"数量" validate:"required,gte=1"`
	Desc bool `json:"desc" label:"降序" validate:"required"`
}
