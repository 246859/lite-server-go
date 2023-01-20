package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"liteserver/global"
	"liteserver/global/code"
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
		claims, err := jwtutils.ParseHs256Jwt(token, "")
		// 如果发生错误的话
		if err != nil {
			c.Abort()
			if errors.Is(err, jwt.ErrTokenExpired) { // token已经过期
				if time.Now().Sub(claims.ExpiresAt.Time) <= global.Config.JwtConfig.AcAllowExpTime() { // 仍然在过期允许时间内
					response.NilBody(c, http.StatusUnauthorized, code.RefreshToken, "token需要刷新")
				} else { // 过期超时
					response.Forbidden(c, code.ExpiredToken, "token已过期")
				}
			} else {
				response.Forbidden(c, code.AccessForbidden, "禁止访问")
			}
		} else {
			c.Set("userClaims", claims)
			c.Next()
		}
	}
}
