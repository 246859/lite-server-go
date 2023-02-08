package public

import (
	"github.com/gin-gonic/gin"
	"liteserver/global"
	"liteserver/service"
	"liteserver/utils/mailutils"
	"liteserver/utils/response"
	"liteserver/utils/validateutils"
	"time"
)

var mailService = service.AppService.MailService

type Mail struct{}

// SendAuthMail
// @Date 2023-02-08 18:06:36
// @Param mail string
// @Method http.MethodGet
// @Description:
func (m *Mail) SendAuthMail(ctx *gin.Context) {
	to := ctx.Query("mail")
	// 参数解析
	if err := ctx.ShouldBind(to); err != nil {
		response.FailWithMsg(ctx, global.I18nRawCN("request.badPrams"))
		return
	}
	if !validateutils.EmailCheck(to) {
		response.FailWithMsg(ctx, global.I18nRawCN("mail.errorFormat"))
		return
	}
	// 发送验证邮件
	mail, err := mailService.SendAuthMail(to)
	if err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}
	// 存入Redis
	global.Redis.Set(ctx, mailutils.RedisMailKey(to, mail.Code), mail.Code, time.Duration(mail.Expire)*time.Minute)
	response.OkWithMsg(ctx, global.I18nRawCN("mail.ok.sendAuth"))
}
