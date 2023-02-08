package sysreq

// Login
// @Date 2023-02-08 21:55:27
// @Description: 登录用户
type Login struct {
	Email    string `json:"email" uri:"email" form:"email" binding:"required,email"`
	Password string `json:"password" uri:"password" form:"password" binding:"required"`
}

// Register
// @Date 2023-02-08 21:55:35
// @Description: 注册用户
type Register struct {
	Email      string `json:"email" binding:"required,email"`
	Nickname   string `json:"nickname" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"rePassword" binding:"required,eqfield=Password" gorm:"-"`
	Ecode      string `json:"ecode" binding:"required"`
}

// ForgetPassword
// @Date 2023-02-08 21:55:42
// @Description: 忘记密码用户
type ForgetPassword struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Ecode    string `json:"ecode" binding:"required"`
}
