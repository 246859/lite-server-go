package system

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"liteserver/global"
	"liteserver/model/sys"
	"liteserver/model/sys/sysrep"
	"liteserver/utils/jwtutils"
	"strings"
	"time"
)

type JwtService struct {
}

// CreateTokenPair
// @Date 2023-02-10 15:48:13
// @Param user sysrep.SystemUser
// @Return sysrep.Jwt
// @Return error
// @Method
// @Description: 根据用户信息创建一对token
func (j JwtService) CreateTokenPair(user sys.SystemUser) (*sysrep.Jwt, error) {
	// 根据用户信息创建Claims
	userClaims := jwtutils.UserClaims{UserId: user.ID, UserUUID: user.Uuid}

	// 创建Refresh Claims
	refreshClaims := jwtutils.CreateJwtClaims(userClaims, jwtutils.JwtCfg.Issuer, jwtutils.JwtCfg.ReExpTime())
	// 创建Access Claims，将refreshClaims的jti作为accessClaims的issuer
	accessClaims := jwtutils.CreateJwtClaims(userClaims, refreshClaims.ID, jwtutils.JwtCfg.AcExpTime())

	// refresh token
	refreshToken, err := jwtutils.CreateHs256Jwt(refreshClaims, jwtutils.JwtCfg.ReSign)
	if err != nil {
		return nil, err
	}
	// access token
	accessToken, err := jwtutils.CreateHs256Jwt(accessClaims, jwtutils.JwtCfg.AcSign)
	if err != nil {
		return nil, err
	}

	return &sysrep.Jwt{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
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
		return "", errors.New(global.I18nRawCN("token.invalid"))
	}

	bearerToken := strings.Split(header, " ")
	if len(bearerToken) < 2 {
		return "", errors.New(global.I18nRawCN("token.invalid"))
	}

	if bearerToken[0] != "Bearer" {
		return "", errors.New(global.I18nRawCN("token.invalidType"))
	}

	if len(bearerToken[1]) == 0 {
		return "", errors.New(global.I18nRawCN("token.invalid"))
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
func (j JwtService) SetJwtToRedis(jwt string, val string) error {
	// redis过期时间是exptime+allowtime
	if err := global.Redis.Set(context.Background(), j.CreateRedisAccessKey(jwt), val, jwtutils.JwtCfg.AcExpTime()+jwtutils.JwtCfg.AcAllowExpTime()).Err(); err != nil {
		return err
	}
	return nil
}

// DelRedisAccessToken
// @Date 2023-02-10 18:01:19
// @Param jwt string
// @Return error
// @Method
// @Description: 删除Redis中的AccessToken
func (j JwtService) DelRedisAccessToken(jwt string) error {
	return global.Redis.Del(context.Background(), j.CreateRedisAccessKey(jwt)).Err()
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

// ParseHs256Token
// @Date 2023-02-10 16:22:46
// @Param token string
// @Param sign string
// @Return jwtutils.Claims
// @Return error
// @Description: 解析Hs256加密的token
func (j JwtService) ParseHs256Token(token string, sign string) (jwtutils.Claims, error) {
	claims, err := jwtutils.ParseHs256Jwt(token, sign)
	switch sign {
	case jwtutils.JwtCfg.AcSign:
		if errors.Is(err, jwt.ErrTokenExpired) {
			// 如果依旧在过期时间内
			if time.Now().Sub(claims.ExpiresAt.Time.Add(jwtutils.JwtCfg.AcAllowExpTime())) <= 0 {
				return claims, jwtutils.ErrJwtNeedToRefresh
			}
		}
	case jwtutils.JwtCfg.ReSign:
		if err != nil {
			return claims, err
		}
	}
	if err != nil {
		return claims, err
	}
	return claims, nil
}

// TokenRefresh
// @Date 2023-02-10 15:32:28
// @Return sysrep.Jwt
// @Return error
// @Description: Token刷新服务
func (j JwtService) TokenRefresh(token sysrep.Jwt) (*sysrep.Jwt, error) {
	accessClaims, err := j.ParseAccessToken(token.Access)
	// 只有access处于需要刷新的状态才有必要刷新token
	if err != nil && !errors.Is(err, jwtutils.ErrJwtNeedToRefresh) {
		return nil, err
	}
	refreshClaims, err := j.ParseRefreshToken(token.Refresh)
	if err != nil {
		return nil, err
	}

	// 对比accessClaims的签发者，如果不是refreshClaims的jti，则说明这一对token没有任何关系
	if accessClaims.Issuer != refreshClaims.ID {
		return nil, errors.New(global.I18nRawCN("token.invalidPair"))
	}

	// 走到这一步说明token合法，两个token也是对应的，可以颁发新的accessToken
	newAccessClaims := jwtutils.CreateJwtClaims(refreshClaims.UserClaims, refreshClaims.ID, jwtutils.JwtCfg.AcExpTime())
	// 创建新的AccessToken
	newAccessToken, err := jwtutils.CreateHs256Jwt(newAccessClaims, jwtutils.JwtCfg.AcSign)
	if err != nil {
		return nil, err
	}
	// 存入Redis
	if err := j.SetJwtToRedis(newAccessToken, refreshClaims.UserUUID); err != nil {
		return nil, err
	}
	return &sysrep.Jwt{
		Access:  newAccessToken,
		Refresh: token.Refresh,
	}, nil
}
