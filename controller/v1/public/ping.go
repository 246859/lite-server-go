package public

import (
	"github.com/gin-gonic/gin"
	"liteserver/utils/response"
)

type Ping struct{}

func (p Ping) Ping(c *gin.Context) {
	response.OkWithMsg(c, "Pong")
}
