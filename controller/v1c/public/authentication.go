package public

import (
	"errors"
	v1 "github.com/246859/lite-server-go/controller/v1c"
	"github.com/246859/lite-server-go/global"
	"github.com/246859/lite-server-go/global/code"
	"github.com/246859/lite-server-go/model/sys/sysrep"
	"github.com/246859/lite-server-go/model/sys/sysreq"
	"github.com/246859/lite-server-go/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var authenService = v1.SystemService.AuthenticationService
var jwtService = v1.SystemService.JwtService

// Authentication
// @Date 2023-01-16 16:04:04
// @Description: 用户认证相关接口
type Authentication struct {
}

// Login
// @Date 2023-01-16 16:04:15
// @Param c *gin.Context
// @Param email string 用户邮箱
// @Param password string 用户密码
// @Method GET
// @Description: 登录接口
func (a Authentication) Login(c *gin.Context) {
	var login sysreq.Login
	// 参数获取
	if err := c.ShouldBind(&login); err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}
	// 调用login服务
	token, err := authenService.Login(&login)
	if err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}
	// 将access token存入redis
	if err := jwtService.SetJwtToRedis(token.Access, login.Email); err != nil {
		response.InternalErrorWithMsg(c, err.Error())
		return
	}
	// 返回结果
	response.OkWithParams(c, code.SuccessLogin, token, global.I18nRawCN("authen.ok.login"))
}

// Register
// @Date 2023-01-16 16:04:48
// @Param c *gin.Context
// @Param registerUser sysreq.Register 用户注册表单
// @Method http.MethodPost
// @Description: 注册接口
func (a Authentication) Register(c *gin.Context) {
	var registerUser sysreq.Register
	// 参数校验
	if err := c.ShouldBind(&registerUser); err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}
	// 调用注册服务
	if err := authenService.Register(&registerUser); err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}
	// 注册成功
	response.OkWithMsg(c, global.I18nRawCN("authen.ok.regiser"))
}

// ForgetPassword
// @Date 2023-02-08 21:24:20
// @Param c *gin.Context
// @Param fgpUser *sysreq.ForgetPassword
// @Method http.MethodPost
// @Description: 忘记密码接口
func (a Authentication) ForgetPassword(c *gin.Context) {
	var fgpUser sysreq.ForgetPassword
	// 参数解析
	if err := c.ShouldBind(&fgpUser); err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}
	if err := authenService.ForgetPassword(&fgpUser); err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}
	response.OkWithMsg(c, global.I18nRawCN("authen.ok.fgp"))
}

// RefreshToken
// @Date 2023-01-16 16:09:00
// @Param c *gin.Context
// @Method http.MethodGet
// @Description: Token刷新接口
func (a Authentication) RefreshToken(c *gin.Context) {
	var oldJwt sysrep.Jwt
	// 先解析参数
	if err := c.ShouldBind(&oldJwt); err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}
	// 随后调用服务
	token, err := jwtService.TokenRefresh(oldJwt)
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			response.Forbidden(c, code.AccessNoLogin, global.I18nRawCN("token.expired"))
		default:
			response.FailWithMsg(c, err.Error())
		}
		return
	}
	// 返回新的token
	response.OkWithParams(c, code.SuccessRefresh, token, global.I18nRawCN("token.refreshOk"))
}
