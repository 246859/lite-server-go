package middleware

import (
	"github.com/246859/lite-server-go/global"
	"github.com/246859/lite-server-go/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// NewRateLimitMiddleware
// @Date 2023-02-10 20:29:24
// @Return gin.HandlerFunc
// @Method
// @Description: 令牌桶限流
func NewRateLimitMiddleware(rate float64, limit int64) gin.HandlerFunc {
	// 创建令牌桶
	bucket := ratelimit.NewBucketWithRate(rate, limit)
	return func(ctx *gin.Context) {
		// 拿出一个令牌，如果拿不出来,则说明可能达到顶峰
		if bucket.TakeAvailable(1) < 1 {
			ctx.Abort()
			response.FailWithMsg(ctx, global.I18nRawCN("response.ratelimit"))
			return
		}
		ctx.Next()
	}
}
