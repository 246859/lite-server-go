package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"liteserver/global"
	"liteserver/global/code"
	"liteserver/model/sys/sysrep"
	"liteserver/utils/jwtutils"
	"liteserver/utils/response"
	"net/http"
	"time"
)

// JwtMiddleWare
// @Date 2023-01-20 22:15:01
// @Return gin.HandlerFunc
// @Method
// @Description: Jwt中间件
func JwtMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求中的Token
		token := c.Request.Header.Get("Authorization")
		// 解析尝试获取claims,这里传入私钥
		claims, err := jwtutils.ParseHs256Jwt(token, global.Config.JwtConfig.AcSign)
		// 从Redis中拿Token
		if len(global.Redis.Get(c, token).Val()) == 0 {
			c.Abort()
			response.Forbidden(c, code.AccessNoLogin, global.I18nRawCN("code.noLogin"))
			return
		}
		// 如果发生错误的话
		if err != nil {
			// 清空后续的调用链，不再继续执行请求
			c.Abort()
			// 根据错误类型进行相应的错误处理
			if errors.Is(err, jwt.ErrTokenExpired) { // token已经过期
				if time.Now().Sub(claims.ExpiresAt.Time) <= global.Config.JwtConfig.AcAllowExpTime() { // 仍然在过期允许时间内
					response.NilBody(c, http.StatusUnauthorized, code.RefreshToken, global.I18nRawCN("code.refreshToken"))
				} else { // 过期超时
					response.Forbidden(c, code.ExpiredToken, global.I18nRawCN("code.expiredToken"))
				}
			} else {
				response.Forbidden(c, code.AccessNoLogin, global.I18nRawCN("code.noLogin"))
			}
		} else {
			// 将用户的信息放入context中
			c.Set(jwtutils.UserClaimsFlag, claims)
			c.Set(jwtutils.UserJwtPayload, sysrep.Jwt{Access: token})
			c.Next()
		}
	}
}
