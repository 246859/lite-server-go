package ginutils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpHandleMethod = func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes

// JudgeMethod
// @Date 2023-01-16 19:03:41
// @Param method string
// @Param group gin.RouterGroup
// @Return httpHandleMethod
// @Description: 根据http请求的类型来返回对应的方法
func JudgeMethod(method string, group *gin.RouterGroup) HttpHandleMethod {
	switch method {
	case http.MethodGet:
		return group.GET
	case http.MethodPost:
		return group.POST
	case http.MethodPut:
		return group.PUT
	case http.MethodDelete:
		return group.DELETE
	case http.MethodOptions:
		return group.OPTIONS
	case http.MethodPatch:
		return group.PATCH
	case http.MethodHead:
		return group.HEAD
	}
	return nil
}
