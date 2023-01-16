package api

import "github.com/gin-gonic/gin"

type LookUpApi = interface{}
type LookUpApiGroup = []LookUpApi

// Api
// @Date 2023-01-15 20:41:17
// @Description: Api接口结构体
type Api struct {
	// 路径
	Path string
	// 接口的方法类型
	Method string
	// 处理函数
	Handler gin.HandlerFunc
	// 中间件
	Middleware []gin.HandlerFunc
}

// ApiGroup
// @Date 2023-01-15 20:41:25
// @Description: Api组结构体
type ApiGroup struct {
	// 路径
	Path string
	// 子分组
	Group LookUpApiGroup
	// 中间件
	Middleware []gin.HandlerFunc
}

func Merge(apiGroups ...LookUpApiGroup) (group LookUpApiGroup) {
	for _, apiGroup := range apiGroups {
		group = append(group, apiGroup...)
	}
	return
}
