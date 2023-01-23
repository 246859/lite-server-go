package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/text/language"
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
		claims, err := jwtutils.ParseHs256Jwt(token, global.Config.JwtConfig.AcSign)
		// 如果发生错误的话
		if err != nil {
			// 清空后续的调用链，不再继续执行请求
			c.Abort()
			// 根据错误类型进行相应的错误处理
			if errors.Is(err, jwt.ErrTokenExpired) { // token已经过期
				if time.Now().Sub(claims.ExpiresAt.Time) <= global.Config.JwtConfig.AcAllowExpTime() { // 仍然在过期允许时间内
					response.NilBody(c, http.StatusUnauthorized, code.RefreshToken, global.I18nLocale.GetWithRaw("code.refreshToken", language.Chinese))
				} else { // 过期超时
					response.Forbidden(c, code.ExpiredToken, global.I18nLocale.GetWithRaw("code.expiredToken", language.Chinese))
				}
			} else {
				response.Forbidden(c, code.AccessNoLogin, global.I18nLocale.GetWithRaw("code.noLogin", language.Chinese))
			}
		} else {
			// 将用户的信息放入context中
			c.Set("userClaims", claims)
			c.Next()
		}
	}
}
