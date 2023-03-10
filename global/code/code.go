package code

// BusinessCode
// @Date: 2023-01-12 22:09:36
// 业务码，非负整数
type BusinessCode = uint32

// @Date: 2023-01-12 22:08:34
const (
	Success = 2000
	// SuccessLogin 登录成功
	SuccessLogin = 2001
	// SuccessRegister 注册成功
	SuccessRegister = 2002
	SuccessLogout   = 2003
	SuccessRefresh  = 2004

	BadOperation = 4000
	// UnAuthorized 未授权
	UnAuthorized = 4010
	// ExpiredToken Token已过期
	ExpiredToken = 4011
	// RefreshToken Token需要刷新
	RefreshToken = 4012
	// AccessForbidden 禁止访问
	AccessForbidden = 4030
	// AccessNoLogin 用户未登录
	AccessNoLogin = 4031
	// AccessNoPerm 没有访问权限
	AccessNoPerm  = 4032
	TokenInvalid  = 4033
	InternalError = 5000
)
