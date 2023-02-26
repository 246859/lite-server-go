package public

import (
	v1 "github.com/246859/lite-server-go/controller/v1c"
	"github.com/246859/lite-server-go/global"
	"github.com/246859/lite-server-go/utils/mailutils"
	"github.com/246859/lite-server-go/utils/responseuils"
	"github.com/246859/lite-server-go/utils/validateutils"
	"github.com/gin-gonic/gin"
	"time"
)

// Mail
// @Date 2023-02-09 19:39:49
// @Description: 公共邮件服务接口
type Mail struct {
}

// SendAuthMail
// @Date 2023-02-08 18:06:36
// @Param email string 邮箱
// @Method http.MethodGet
// @Description:
func (m *Mail) SendAuthMail(ctx *gin.Context) {
	to := ctx.Query("email")
	// 参数解析
	if err := ctx.ShouldBind(to); err != nil {
		responseuils.FailWithMsg(ctx, global.I18nRawCN("request.badPrams"))
		return
	}
	if err := validateutils.EmailCheck(to); err != nil {
		responseuils.FailWithMsg(ctx, global.I18nRawCN("mail.errorFormat"))
		return
	}
	// 发送验证邮件
	mail, err := v1.MailSevice.SendAuthMail(to)
	if err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
		return
	}
	// 存入Redis
	global.Redis.Set(ctx, mailutils.RedisMailKey(to, mail.Code), mail.Code, time.Duration(mail.Expire)*time.Minute)
	responseuils.OkWithMsg(ctx, global.I18nRawCN("mail.ok.sendAuth"))
}
