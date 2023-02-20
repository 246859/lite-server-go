package v1r

import "liteserver/controller"

var (
	AppController    = controller.ControllerGroup
	PublicController = AppController.Public
	SystemController = AppController.System
)

// 公共接口
var (
	// Ping 连通测试接口
	Ping = PublicController.Ping
	// AutenController 用户认证接口
	AutenController = PublicController.Authentication
	// MailController 邮箱接口
	MailController = PublicController.Mail
	// UserSimpleController 用户简单接口
	UserSimpleController = PublicController.UserSimpleController
	// ArticleController 文章接口
	ArticleController = PublicController.ArticleController
)

// 私有接口
var (
	// Pong 私有连通测试接口
	Pong = SystemController.Pong
	// ArticleModifyController 文章修改接口
	ArticleModifyController = SystemController.ArticleModifyController
	// UserController 用户接口
	UserController = SystemController.UserController
	// AuthorizationController 授权接口
	AuthorizationController = SystemController.Authorization

	FileController = SystemController.FileController
)
