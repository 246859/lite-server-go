package system

import (
	"github.com/246859/lite-server-go/utils/response"
	"github.com/gin-gonic/gin"
)

type Pong struct {
}

// Pong
// @Date 2023-02-09 20:35:27
// @Param ctx *gin.Context
// @Method http.MethodGet
// @Description: 私有接口连通测试接口
func (p Pong) Pong(ctx *gin.Context) {
	response.OkWithMsg(ctx, "ping")
}
