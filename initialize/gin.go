package initialize

import (
	"github.com/246859/lite-server-go/config"
	"github.com/246859/lite-server-go/middleware"
	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-redis/redis/v8"

	"net/http"
	"time"
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
// @Description: 安装Gin全局插件
func installPlugins(engine *gin.Engine, cfg *config.ServerConfig, redisClient *redis.Client) {
	engine.Use(
		// 限流中间件
		middleware.NewRateLimitMiddleware(cfg.Rate, cfg.Limit),
		// 日志组件
		middleware.ZapLogger(),
		// 日志错误记录组件
		middleware.ZapRecovery(),
		// 缓存中间件
		cache.CacheByRequestURI(persist.NewRedisStore(redisClient), time.Duration(cfg.Cache)*time.Second))

	// 字段验证翻译器
	binding.Validator = middleware.UniverseValidateTranslator()
}

// installRouter
// @Date 2023-01-11 16:39:10
// @Param engine *gin.Engine
// @Description: 安装Gin路由
func installRouter(engine *gin.Engine, cfg *config.ServerConfig) {
	InitRouter(engine, cfg)
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

func InitHttpServer(cfg *config.ServerConfig, client *redis.Client) *http.Server {
	engine := newEngine(cfg)
	installPlugins(engine, cfg, client)
	installRouter(engine, cfg)
	return configHttpServer(engine, cfg)
}
