package route

import (
	"github.com/gin-gonic/gin"
)

// Api
// @Date 2023-01-16 21:36:36
// @Description: 表示一个对应的接口
type Api struct {
	Mds *MiddleWareRoute
	// 路径
	Path string
	// 接口的方法类型
	Method string
	// 处理函数
	Handler gin.HandlerFunc
}

// ApiGroup
// @Date 2023-01-16 21:37:36
// @Description: 表示一个对应的接口分组
type ApiGroup struct {
	Mds *MiddleWareRoute
	// 是否构成URL
	IsUrl bool
	// 路径
	Path string
	// 该分组下的一组平行路由
	Router Router
	// 该分组下的子路由分组，当Router和Group两者都不为空时，Group的优先级大于Router
	Group RouterGroup
}

type RouterMap = map[string]*Api

type RouterGroupMap = map[string]*ApiGroup

// RouterGroup
// @Date 2023-01-16 21:40:22
// @Description: 路由分组接口，一个路由分组下可以包含多个API，也可以包含多个分组，但都是以 ApiGroup 的形式呈现
type RouterGroup interface {
	InitGroup() RouterGroupMap
}

// Router
// @Date 2023-01-16 21:40:32
// @Description: 路由接口，一个路由接口下可以包含多个API
type Router interface {
	InitRouter() RouterMap
}
