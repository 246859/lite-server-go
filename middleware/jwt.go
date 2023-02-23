package middleware

import (
	"errors"
	"github.com/246859/lite-server-go/global"
	"github.com/246859/lite-server-go/global/code"
	"github.com/246859/lite-server-go/model/response"
	"github.com/246859/lite-server-go/service"
	"github.com/246859/lite-server-go/utils/jwtutils"
	"github.com/246859/lite-server-go/utils/responseuils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"net/http"
)

var jwtService = service.AppService.JwtService

// JwtMiddleWare
// @Date 2023-01-20 22:15:01
// @Return gin.HandlerFunc
// @Method
// @Description: Jwt中间件
func JwtMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取jwt
		accessToken, err := jwtService.BearerTokenFromHeader(c)
		if err != nil {
			c.Abort()
			responseuils.Forbidden(c, code.TokenInvalid, err.Error())
			return
		}

		// 判断该jwt是否存在于redis中
		if !jwtService.CheckJwtFromRedis(c, accessToken) {
			c.Abort()
			responseuils.Forbidden(c, code.AccessNoLogin, global.I18nRawCN("code.noLogin"))
			return
		}
		// 解析access token
		claims, err := jwtService.ParseAccessToken(accessToken)
		if err != nil {
			c.Abort()
			switch {
			case errors.Is(err, jwtutils.ErrJwtNeedToRefresh):
				responseuils.NilBody(c, http.StatusUnauthorized, code.RefreshToken, global.I18nRawCN("code.refreshToken"))
			case errors.Is(err, jwt.ErrTokenExpired):
				responseuils.Forbidden(c, code.ExpiredToken, global.I18nRawCN("code.expiredToken"))
			default:
				responseuils.Forbidden(c, code.AccessNoLogin, global.I18nRawCN("code.noLogin"))
			}
		} else {
			// 将用户的信息放入context中
			c.Set(jwtutils.UserClaimsFlag, claims)
			c.Set(jwtutils.UserJwtPayload, response.Jwt{Access: accessToken})
			c.Next()
		}
	}
}
