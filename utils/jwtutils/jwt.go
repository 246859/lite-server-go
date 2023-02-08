package jwtutils

import (
	"github.com/golang-jwt/jwt/v4"
	"liteserver/config"
	"liteserver/utils/uuidtool"
	"time"
)

const UserClaimsFlag = "userClaims"
const UserJwtPayload = "userPayload"

var JwtCfg *config.JwtConfig

// Claims
// @Date 2023-01-20 20:10:39
// @Description: 自定义Token声明
type Claims struct {
	UserClaims
	jwt.RegisteredClaims
}

type UserClaims struct {
	UserId   uint   `json:"userId"`
	UserUUID string `json:"userUUID"`
}

// CreateHs256Jwt
// @Date 2023-01-20 20:35:38
// @Param signKey string
// @Param userId string
// @Param issuer string
// @Param expired time.Duration
// @Return string
// @Return error
// @Method
// @Description: 使用HMCA256对称加密算法创建一个token
func CreateHs256Jwt(userClaims UserClaims, secret string, issuer string, expired time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserClaims: userClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expired)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuidtool.NewUUIDv5(),
		},
	})
	jwtString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return jwtString, nil
}

// CreateAccessToken
// @Date 2023-02-07 20:29:48
// @Param userClaims UserClaims
// @Return string
// @Return error
// @Method
// @Description: 创建AccessToken
func CreateAccessToken(userClaims UserClaims) (string, error) {
	return CreateHs256Jwt(userClaims, JwtCfg.AcSign, JwtCfg.Issuer, JwtCfg.AcExpTime())
}

// CreateRefreshToken
// @Date 2023-02-07 20:29:56
// @Param userClaims UserClaims
// @Return string
// @Return error
// @Method
// @Description: 创建RefreshToken
func CreateRefreshToken(userClaims UserClaims) (string, error) {
	return CreateHs256Jwt(userClaims, JwtCfg.ReSign, JwtCfg.Issuer, JwtCfg.ReExpTime())
}

func SetConfig(cfg *config.JwtConfig) {
	JwtCfg = cfg
}

// ParseHs256Jwt
// @Date 2023-01-20 22:03:46
// @Param jwtStr string
// @Param secret string
// @Return Claims
// @Return error
// @Method
// @Description: 解析一个HS256签名的JWT
func ParseHs256Jwt(jwtStr string, secret string) (Claims, error) {
	token, err := jwt.ParseWithClaims(jwtStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}, jwt.WithValidMethods([]string{"HS256"}), jwt.WithJSONNumber())

	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return *claims, nil
		}
	}

	return Claims{}, err
}
