package request

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
	Email      string `json:"email" binding:"required,email" label:"用户邮箱" `
	Nickname   string `json:"nickname" binding:"required" label:"用户昵称"`
	Password   string `json:"password" binding:"required" label:"用户密码"`
	RePassword string `json:"rePassword" binding:"required,eqfield=Password" label:"二次密码" gorm:"-"`
	Ecode      string `json:"ecode" binding:"required" label:"邮箱验证码"`
}

// ForgetPassword
// @Date 2023-02-08 21:55:42
// @Description: 忘记密码用户
type ForgetPassword struct {
	Email    string `json:"email" binding:"required,email" label:"用户邮箱" `
	Password string `json:"password" binding:"required" label:"用户密码"`
	Ecode    string `json:"ecode" binding:"required" label:"邮箱验证码"`
}

// UpdateUser
// @Date 2023-02-25 19:23:40
// @Description: 更新用户结构体
type UpdateUser struct {
	Avatar      string `json:"avatar" label:"头像" binding:"required"`
	Nickname    string `json:"nickname" label:"昵称" binding:"required"`
	Password    string `json:"password" label:"密码" binding:"required"`
	Description string `json:"description" label:"自我描述" binding:"required"`
}
