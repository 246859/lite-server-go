package public

import (
	"github.com/gin-gonic/gin"
	"liteserver/utils/response"
)

type Ping struct{}

// Ping
// @Date 2023-02-09 19:40:11
// @Param c *gin.Context
// @Method http.MethodGet
// @Description: 公共连通测试接口
func (p Ping) Ping(c *gin.Context) {
	response.OkWithMsg(c, "Pong")
}
