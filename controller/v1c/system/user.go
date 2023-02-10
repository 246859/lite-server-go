package system

import (
	"github.com/gin-gonic/gin"
	"liteserver/controller/v1c"
	"liteserver/global"
	"liteserver/model/sys/sysrep"
	"liteserver/utils/jwtutils"
	"liteserver/utils/response"
)

// UserController
// @Date 2023-02-09 19:39:12
// @Description: 用户信息操作相关接口
type UserController struct {
}

// UserInfo
// @Date 2023-02-09 19:38:25
// @Param uuid string 用户标识ID
// @Method http.MethodGet
// @Description: 获取用户信息
func (u UserController) UserInfo(ctx *gin.Context) {

}

// UserList
// @Date 2023-02-09 19:41:09
// @Param ctx *gin.Context
// @Method http.MethodGet
// @Description: 分页获取用户列表
func (u UserController) UserList(ctx *gin.Context) {

}

// UpdateUserInfo
// @Date 2023-02-09 19:44:34
// @Method http.MethodGet
// @Description: 更新用户个人信息
func (u UserController) UpdateUserInfo(ctx *gin.Context) {

}

// DeleteUser
// @Date 2023-02-09 20:02:28
// @Param ctx *gin.Context
// @Method
// @Description: 注销用户接口
func (u UserController) DeleteUser(ctx *gin.Context) {

}

// Logout
// @Date 2023-02-09 19:46:25
// @Param ctx *gin.Context
// @Method
// @Description: 用户登出接口
func (u UserController) Logout(ctx *gin.Context) {
	fail := false
	// 从Redis中获取token
	value, exists := ctx.Get(jwtutils.UserJwtPayload)
	// 删除Redis值
	if jwtObj, ok := value.(sysrep.Jwt); ok && exists {
		if v1c.SystemService.JwtService.DelRedisAccessToken(jwtObj.Access) != nil {
			fail = true
		}
	} else {
		fail = true
	}
	if fail {
		response.FailWithMsg(ctx, global.I18nRawCN("authen.fail.logout"))
	} else {
		response.OkWithMsg(ctx, global.I18nRawCN("authen.ok.logout"))
	}
}

// ChangePassword
// @Date 2023-02-09 19:47:08
// @Param ctx *gin.Context
// @Method
// @Description: 用户修改密码接口
func (u UserController) ChangePassword(ctx *gin.Context) {

}
