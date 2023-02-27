package system

import (
	"github.com/246859/lite-server-go/controller/v1c"
	"github.com/246859/lite-server-go/global"
	"github.com/246859/lite-server-go/global/code"
	"github.com/246859/lite-server-go/model/request"
	"github.com/246859/lite-server-go/model/response"
	"github.com/246859/lite-server-go/utils/jwtutils"
	"github.com/246859/lite-server-go/utils/responseuils"
	"github.com/gin-gonic/gin"
)

// UserController
// @Date 2023-02-09 19:39:12
// @Description: 用户信息操作相关接口
type UserController struct {
}

// UserBasicInfo
// @Date 2023-02-26 21:15:57
// @Method http.MethodGet
// @Param id int
// @Url /user/basic
// @Description:
func (u UserController) UserBasicInfo(ctx *gin.Context) {
	var idw request.IdWrap
	if err := ctx.ShouldBind(&idw); err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
		return
	}

	info, err := v1c.UserService.GetUserBasicInfo(idw.Id)

	if err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
	} else {
		responseuils.OkWithMsgAndData(ctx, info, global.I18nRawCN("user.info.ok"))
	}
}

// UserSimpleInfo
// @Date 2023-02-26 21:20:21
// @Method http.MethodGet
// @Param id int
// @Url /user/simple
// @Description:
func (u UserController) UserSimpleInfo(ctx *gin.Context) {
	var idw request.IdWrap
	if err := ctx.ShouldBind(&idw); err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
		return
	}

	info, err := v1c.UserService.GetUserSimpleInfo(idw.Id)

	if err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
	} else {
		responseuils.OkWithMsgAndData(ctx, info, global.I18nRawCN("user.info.ok"))
	}
}

// ListUserBasicInfo
// @Date 2023-02-26 21:19:18
// @Method http.MethodGet
// @Param pageInfo request.PageInfo
// @url /user/basics
// @Description:
func (u UserController) ListUserBasicInfo(ctx *gin.Context) {
	var pageInfo request.PageInfo
	if err := ctx.ShouldBindUri(pageInfo); err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
	}

	list, err := v1c.UserService.GetUserBasicInfoList(pageInfo)

	if err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
	} else {
		responseuils.OkWithMsgAndData(ctx, list, global.I18nRawCN("user.list.ok"))
	}
}

// ListUserSimpleInfo
// @Date 2023-02-26 21:19:50
// @Method http.MethodGet
// @Param pageInfo request.PageInfo
// @Url /user/simples
// @Description: 用户简单信息列表接口
func (u UserController) ListUserSimpleInfo(ctx *gin.Context) {
	var pageInfo request.PageInfo
	if err := ctx.ShouldBindUri(pageInfo); err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
	}

	list, err := v1c.UserService.GetUserSimpleInfoList(pageInfo)

	if err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
	} else {
		responseuils.OkWithMsgAndData(ctx, list, global.I18nRawCN("user.list.ok"))
	}
}

// UpdateUserInfo
// @Date 2023-02-09 19:44:34
// @Method http.MethodGet
// @Description: 用户个人信息更新接口
func (u UserController) UpdateUserInfo(ctx *gin.Context) {
	var updateUser request.UpdateUser
	if err := ctx.ShouldBind(&updateUser); err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
		return
	}

	claims, err := jwtutils.ToJwtClaims(ctx)
	if err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
		return
	}
	err = v1c.UserService.UpdateUserInfo(claims.UserId, updateUser)

	if err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
	} else {
		responseuils.OkWithMsg(ctx, global.I18nRawCN("user.update.ok"))
	}
}

// DeleteUser
// @Date 2023-02-09 20:02:28
// @Param ctx *gin.Context
// @Method http.MethodDelete
// @Param id int
// @Description: 删除用户接口
func (u UserController) DeleteUser(ctx *gin.Context) {
	var idw request.IdWrap
	if err := ctx.ShouldBind(&idw); err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
		return
	}

	err := v1c.UserService.DeleteUser(idw.Id)

	if err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
	} else {
		responseuils.OkWithMsg(ctx, global.I18nRawCN("user.delete.ok"))
	}
}

// Logout
// @Date 2023-02-09 19:46:25
// @Param ctx *gin.Context
// @Method http.MethodDelete
// @Description: 用户登出接口
func (u UserController) Logout(ctx *gin.Context) {
	fail := false
	// 从Redis中获取token
	value, exists := ctx.Get(jwtutils.UserJwtPayload)
	// 删除Redis值
	if jwtObj, ok := value.(response.Jwt); ok && exists {
		if v1c.SystemService.JwtService.DelRedisAccessToken(jwtObj.Access) != nil {
			fail = true
		}
	} else {
		fail = true
	}

	if fail {
		responseuils.FailWithMsg(ctx, global.I18nRawCN("authen.fail.logout"))
	} else {
		responseuils.OkWithParams(ctx, code.SuccessLogout, nil, global.I18nRawCN("authen.ok.logout"))
	}
}

// ChangePassword
// @Date 2023-02-09 19:47:08
// @Param ctx *gin.Context
// @Method
// @Description: 用户修改密码接口
func (u UserController) ChangePassword(ctx *gin.Context) {

}
