package system

import "github.com/gin-gonic/gin"

// Authentication
// @Date 2023-01-16 16:04:04
// @Description: 用户认证相关的接口
type Authentication struct {
}

// Login
// @Date 2023-01-16 16:04:15
// @Param c *gin.Context
// @Description: 登录接口
func (a Authentication) Login(c *gin.Context) {

}

// Logout
// @Date 2023-01-16 16:04:37
// @Param c *gin.Context
// @Description: 登出接口
func (a Authentication) Logout(c *gin.Context) {

}

// Register
// @Date 2023-01-16 16:04:48
// @Param c *gin.Context
// @Description: 注册接口

func (a Authentication) Register(c *gin.Context) {

}

// ForgetPassword
// @Date 2023-01-16 16:05:16
// @Param c *gin.Context
// @Description: 忘记密码接口
func (a Authentication) ForgetPassword(c *gin.Context) {

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
