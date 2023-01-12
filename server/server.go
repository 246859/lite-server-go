package server

import (
	"context"
	"errors"
	"flag"
	"liteserver/utils"
	"log"
	"net/http"

	"github.com/fsnotify/fsnotify"
	"go.uber.org/zap"
	"liteserver/config"
	"liteserver/global"
	"liteserver/initialize"
)

var (
	defaultConfigPath = "./config.yml"

	errLoadEnv    = errors.New("应用环境加载异常")
	errLoadConfig = errors.New("应用配置加载异常")
)

type Server struct {
	ConfigPath string
	server     *http.Server
}

// loadConfig
// @Date: 2023-01-08 22:17:12
// @Description: 刷新应用配置
// @Receiver: s *Server
// @Return: error
func (s *Server) loadConfig() error {
	// 读取配置
	viper, err := config.ReadConfig(s.ConfigPath)
	if err != nil {
		return err
	}
	global.Viper = viper
	// 加载配置
	serverConfig, err := config.RefreshConfig(global.Viper)
	if err != nil {
		return err
	}

	global.Config = serverConfig
	return nil
}

// loadEnv
// @Date: 2023-01-08 22:17:25
// @Description: 刷新应用环境
// @Receiver: s *Server
// @Return: error
func (s *Server) loadEnv() error {
	if global.Config == nil {
		return errLoadConfig
	}
	// 初始化日志
	global.Logger = initialize.InitZap(global.Config.ZapConfig)
	zap.L().Info("应用日志系统初始化完毕!")
	// 初始化国际化语言信息
	global.I18nLocale = initialize.InitI18nInfo(global.Config.I18nConfig)
	// 初始化Redis连接
	global.Redis = initialize.InitRedis(global.Config.RedisConfig)
	// 初始化GORM和数据库
	global.GormDBGroup = initialize.InitGorm(global.Config.DataBaseConfig)
	// 初始化Http服务器
	s.server = initialize.InitHttpServer(global.Config.ServerConfig)
	return nil
}

// runHttpServer
// @Date 2023-01-11 17:47:36
// @Description: 运行Http服务器
func (s *Server) runHttpServer() {
	defer s.ShutDown()
	err := s.server.ListenAndServe()
	if err != nil {
		zap.L().Error("应用启动失败，即将关闭", zap.Error(err))
	}
}

// Run
// @Date: 2023-01-08 22:17:38
// @Description: 应用启动
// @Receiver: s *Server
func (s *Server) Run() {
	utils.LogBanner()
	// 加载配置文件
	err := s.loadConfig()
	if err != nil {
		log.Panicln("配置文件加载失败:", err)
	}

	// 加载应用环境
	err = s.loadEnv()
	if err != nil {
		log.Panicln("应用环境初始化失败:", err)
	}
	s.refresh()
	s.runHttpServer()
}

// ShutDown
// @Date: 2023-01-09 23:28:26
// @Description: 优雅的关闭应用
// @Receiver: s *Server
func (s *Server) ShutDown() {
	s.beforeShutDown()
	err := s.server.Shutdown(context.Background())
	if err != nil {
		zap.L().Fatal("Http服务器无法正常关闭，应用程序将立即停止", zap.Error(err))
	}
}

func (s *Server) beforeShutDown() {
	zap.L().Info("应用正常关闭")
	// TODO 这里做一些在关闭服务器之前要做的事情
}

func (s *Server) refresh() {
	global.Viper.WatchConfig()
	global.Viper.OnConfigChange(func(in fsnotify.Event) {
		// 加载应用环境
		err := s.loadEnv()
		if err != nil {
			log.Panicln("应用环境刷新失败:", err)
		}
	})
}

// RunWithFlag
// @Date: 2023-01-08 22:22:37
// @Description: 使用命令行标志输入配置文件路径
// @Receiver: s *Server
func (s *Server) RunWithFlag() {
	configPath := flag.String("cfg", defaultConfigPath, "服务器配置文件")
	flag.Parse()
	s.ConfigPath = *configPath
	s.Run()
}

// RunWithDefault
// @Date: 2023-01-08 22:22:28
// @Description: 使用默认的配置文件路径
// @Receiver: s *Server
func (s *Server) RunWithDefault() {
	s.ConfigPath = defaultConfigPath
	s.Run()
}
