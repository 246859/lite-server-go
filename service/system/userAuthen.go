package system

import (
	"context"
	"errors"
	"github.com/246859/lite-server-go/global"
	"github.com/246859/lite-server-go/model"
	"github.com/246859/lite-server-go/model/request"
	"github.com/246859/lite-server-go/model/response"
	"github.com/246859/lite-server-go/utils/ginutils"
	"github.com/246859/lite-server-go/utils/mailutils"
	"github.com/246859/lite-server-go/utils/uuidtool"
)

type AuthenticationService struct{}

// Login
// @Date 2023-02-06 18:08:43
// @Param user *sys.SystemUser
// @Return userInfo *sys.SystemUser
// @Return err error
// @Description: 登录服务
func (a *AuthenticationService) Login(loginuser *request.Login) (userInfo *response.Jwt, err error) {
	var sysUser model.SystemUser
	// 查询数据库中是否含有该对象
	global.DB().Model(model.SystemUser{}).Where("email = ?", loginuser.Email).First(&sysUser)
	// 如果用户不存在
	if len(sysUser.Email) == 0 {
		return nil, errors.New(global.I18nRawCN("authen.userNotFound"))
	}
	// 验证密码是否相等
	if sysUser.Password != ginutils.Sha1(loginuser.Password) {
		return nil, errors.New(global.I18nRawCN("authen.passwordError"))
	}

	// 创建token对
	token, err := new(JwtService).CreateTokenPair(sysUser)

	if err != nil {
		return nil, err
	}
	return token, nil
}

// Register
// @Date 2023-02-08 20:25:06
// @Param user *sysreq.Register
// @Return err error
// @Description: 注册服务
func (a *AuthenticationService) Register(regiUser *request.Register) (err error) {
	var sysUser model.SystemUser
	// 查询是否已存在该用户
	global.DB().Model(sysUser).Where("email=?", regiUser.Email).First(&sysUser)
	// 如果用户已经存在
	if len(sysUser.Email) > 0 {
		return errors.New(global.I18nRawCN("authen.userExist"))
	}
	if regiUser.Password != regiUser.RePassword {
		return errors.New(global.I18nRawCN("authen.fail.notSamePasd"))
	}
	// 获取键值
	redisKey := mailutils.RedisMailKey(regiUser.Email, regiUser.Ecode)
	// 从Redis中拿验证码
	val := global.Redis.Del(context.Background(), redisKey).Val()
	// 如果验证码不正确
	if val == 0 {
		return errors.New(global.I18nRawCN("authem.fail.errcode"))
	}

	// 复制属性
	sysUser.Nickname = regiUser.Nickname
	sysUser.Uuid = uuidtool.NewUUIDv5()
	sysUser.Email = regiUser.Email
	sysUser.Password = ginutils.Sha1(regiUser.RePassword)
	// 插入数据库
	if err := global.DB().Model(sysUser).Create(&sysUser).Error; err != nil {
		return err
	}
	// 成功过后删除Redis中的key
	global.Redis.Del(context.Background(), redisKey)
	return nil
}

func (a *AuthenticationService) ChangePassword(user *model.SystemUser) (err error) {
	return nil
}

// ForgetPassword
// @Date 2023-02-08 21:29:46
// @Param user *sysreq.ForgetPassword
// @Return err error
// @Method
// @Description: 忘记密码服务
func (a *AuthenticationService) ForgetPassword(fpgUser *request.ForgetPassword) (err error) {
	var sysUser model.SystemUser
	// 首先根据邮箱查找用户
	global.DB().Model(sysUser).Where("email=?", fpgUser.Email).First(&sysUser)
	// 如果用户不存在，那么就返回错误
	if len(sysUser.Email) == 0 {
		return errors.New(global.I18nRawCN("authen.userNotFound"))
	}
	// 如果验证码不相等
	redisKey := mailutils.RedisMailKey(fpgUser.Email, fpgUser.Ecode)
	if global.Redis.Del(context.Background(), redisKey).Val() == 0 {
		return errors.New(global.I18nRawCN("authem.fail.errcode"))
	}
	// 最后修改密码
	if err := global.DB().Model(sysUser).Update("password", ginutils.Sha1(fpgUser.Password)).Error; err != nil {
		return err
	}
	// 成功过后删除Redis中的key
	global.Redis.Del(context.Background(), redisKey)
	return nil
}
