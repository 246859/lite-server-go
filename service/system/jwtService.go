package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"liteserver/global"
	"liteserver/utils/jwtutils"
	"strings"
	"time"
)

type JwtService struct {
}

// BearerTokenFromHeader
// @Date 2023-02-09 22:18:58
// @Param c *gin.Context
// @Return string
// @Return error
// @Method
// @Description: 从请求头中获取Bearer token
func (j JwtService) BearerTokenFromHeader(c *gin.Context) (string, error) {
	header := c.Request.Header.Get("Authorization")
	if len(header) == 0 {
		return "", errors.New("access-token不存在")
	}

	bearerToken := strings.Split(header, " ")
	if len(bearerToken) < 2 {
		return "", errors.New("access-token不存在")
	}

	if bearerToken[0] != "Bearer" {
		return "", errors.New("不是Bearer类型Token")
	}

	if len(bearerToken[1]) == 0 {
		return "", errors.New("token不存在")
	}

	return bearerToken[1], nil
}

// CheckJwtFromRedis
// @Date 2023-02-09 22:20:37
// @Param c gin.Context
// @Param jwt string
// @Return bool
// @Method
// @Description: 检查Redis中的jwt是否存在
func (j JwtService) CheckJwtFromRedis(c *gin.Context, jwt string) bool {
	return len(global.Redis.Get(c, j.CreateRedisAccessKey(jwt)).Val()) != 0
}

// SetJwtToRedis
// @Date 2023-02-09 22:25:42
// @Param c *gin.Context
// @Param jwt string
// @Method
// @Description: 将Jwt存入Redis中
func (j JwtService) SetJwtToRedis(c *gin.Context, jwt string, val string) error {
	// redis过期时间是exptime+allowtime
	if err := global.Redis.Set(c, j.CreateRedisAccessKey(jwt), val, jwtutils.JwtCfg.AcExpTime()+jwtutils.JwtCfg.AcAllowExpTime()).Err(); err != nil {
		return err
	}
	return nil
}

// CreateRedisAccessKey
// @Date 2023-02-09 22:28:38
// @Description: 创建AccessToken在Redis的key
func (j JwtService) CreateRedisAccessKey(jwt string) string {
	return "user-access-token-" + jwt
}

// ParseAccessToken
// @Date 2023-02-09 22:51:06
// @Param token string
// @Return jwtutils.Claims
// @Return error
// @Method
// @Description: 解析Access token
func (j JwtService) ParseAccessToken(token string) (jwtutils.Claims, error) {
	return j.ParseHs256Token(token, jwtutils.JwtCfg.AcSign)
}

// ParseRefreshToken
// @Date 2023-02-09 22:51:19
// @Param token string
// @Return jwtutils.Claims
// @Return error
// @Method
// @Description: 解析Refresh Token
func (j JwtService) ParseRefreshToken(token string) (jwtutils.Claims, error) {
	return j.ParseHs256Token(token, jwtutils.JwtCfg.ReSign)
}

func (j JwtService) ParseHs256Token(token string, sign string) (jwtutils.Claims, error) {
	claims, err := jwtutils.ParseHs256Jwt(token, sign)
	// 如果过期了分两种情况，完全过期和还在允许时间内
	if errors.Is(err, jwt.ErrTokenExpired) {
		// 如果依旧在过期时间内
		if time.Now().Sub(claims.ExpiresAt.Time.Add(jwtutils.JwtCfg.AcAllowExpTime())) <= 0 {
			return claims, jwtutils.ErrJwtNeedToRefresh
		}
	}
	if err != nil {
		return claims, err
	}
	return claims, nil
}
