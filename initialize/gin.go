package initialize

import (
	"github.com/gin-gonic/gin"
	"liteserver/config"
	"liteserver/middleware"
	"net/http"
)

// newEngine
// @Date 2023-01-11 17:04:43
// @Param cfg config.ServerConfig
// @Return http.Handler
// @Description:  新建一个gin引擎
func newEngine(cfg *config.ServerConfig) *gin.Engine {
	engine := gin.New()
	// TODO 这里以后可以做一些初始化的工作
	return engine
}

// installPlugins
// @Date 2023-01-11 16:35:02
// @Param engine *gin.Engine
// @Description: 安装Gin插件
func installPlugins(engine *gin.Engine) {
	engine.Use(
		middleware.ZapLogger(),
		middleware.ZapRecovery())
}

// installRouter
// @Date 2023-01-11 16:39:10
// @Param engine *gin.Engine
// @Description: 安装Gin路由
func installRouter(engine *gin.Engine) {
	// TODO 以后这里要初始化Router
}

// configHttpServer
// @Date 2023-01-11 16:40:13
// @Param engine *gin.Engine
// @Param cfg *config.ServerConfig
// @Return *http.Server
// @Description: 配置Http服务器
func configHttpServer(engine *gin.Engine, cfg *config.ServerConfig) *http.Server {
	engine.MaxMultipartMemory = cfg.MultipartMemory
	gin.SetMode(cfg.Mode)
	server := &http.Server{
		Handler:           engine,
		Addr:              cfg.Addr(),
		ReadTimeout:       cfg.ReadTimeOut(),
		WriteTimeout:      cfg.WriteTimeOut(),
		IdleTimeout:       cfg.IdleTimeOut(),
		ReadHeaderTimeout: cfg.ReadHeaderTimeOut(),
		MaxHeaderBytes:    cfg.MaxHeaderBytes,
	}
	return server
}

func InitHttpServer(cfg *config.ServerConfig) *http.Server {
	engine := newEngine(cfg)
	installPlugins(engine)
	installRouter(engine)
	return configHttpServer(engine, cfg)
}
