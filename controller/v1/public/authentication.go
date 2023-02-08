package public

import (
	"github.com/gin-gonic/gin"
	"liteserver/global"
	"liteserver/model/sys/sysreq"
	"liteserver/service"
	"liteserver/utils/jwtutils"
	"liteserver/utils/response"
)

var authenService = service.AppService.AuthenticationService

// Authentication
// @Date 2023-01-16 16:04:04
// @Description: 用户认证相关的接口
type Authentication struct {
}

// Login
// @Date 2023-01-16 16:04:15
// @Param c *gin.Context
// @Param username string
// @Param password string
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
	// 将生成的Token存入Redis
	global.Redis.Set(c, token.Access, login.Email, jwtutils.JwtCfg.AcExpTime())
	global.Redis.Set(c, token.Refresh, login.Email, jwtutils.JwtCfg.ReExpTime())
	// 返回结果
	if err == nil {
		response.OkWithMsgAndData(c, token, global.I18nRawCN("authen.ok.login"))
	} else {
		response.FailWithMsg(c, err.Error())
	}
}

// Logout
// @Date 2023-01-16 16:04:37
// @Param c *gin.Context
// @Description: 登出接口
func (a Authentication) Logout(c *gin.Context) {
	fail := false
	// 从Redis中获取token
	value, exists := c.Get(jwtutils.UserJwtPayload)
	// 删除Redis值
	if v, ok := value.(string); ok && exists {
		if global.Redis.Del(c, v).Val() == 0 {
			fail = true
		}
	} else {
		fail = true
	}
	if fail {
		response.FailWithMsg(c, global.I18nRawCN("authen.fail.logout"))
	} else {
		response.OkWithMsg(c, global.I18nRawCN("authen.ok.logout"))
	}
}

// Register
// @Date 2023-01-16 16:04:48
// @Param c *gin.Context
// @Param registerUser sysreq.Register
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

// ChangePassword
// @Date 2023-01-16 16:08:53
// @Param c *gin.Context
// @Description: 修改密码接口
func (a Authentication) ChangePassword(c *gin.Context) {

}

// RefreshToken
// @Date 2023-01-16 16:09:00
// @Param c *gin.Context
// @Description: Token刷新接口
func (a Authentication) RefreshToken(c *gin.Context) {

}
