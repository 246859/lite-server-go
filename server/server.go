package server

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"go.uber.org/zap"
	"liteserver/config"
	"liteserver/global"
	"liteserver/initialize"
	"liteserver/model"
	"liteserver/resource"
	"liteserver/utils"
	"liteserver/utils/fileutils"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var (
	errLoadEnv    = errors.New("应用环境加载异常")
	errLoadConfig = errors.New("应用配置加载异常")

	defaultConfig           = "template/defaultConfig.yml"
	defaultTargetConfigPath = "config.yml"
)

func init() {
	global.WorkDir = fileutils.GetCurrentPath()
}

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
	// 输出日志，说明日志系统初始化完毕
	zap.L().Info("应用日志系统初始化完毕!")
	// 初始化国际化语言信息
	global.I18nLocale = initialize.InitI18nInfo(global.Config.I18nConfig)
	zap.L().Info("应用多语言配置初始化完毕!")
	// 初始化Redis连接
	global.Redis = initialize.InitRedis(global.Config.RedisConfig)
	zap.L().Info("应用Redis连接成功!")
	// 初始化GORM和数据库
	global.GormDBGroup = initialize.InitGorm(global.Config.DataBaseConfig)
	zap.L().Info("应用GORM初始化成功!")
	// 初始化JWT配置
	initialize.InitJwtUtils(global.Config.JwtConfig)
	// 初始化表
	initialize.IniTables(model.ModelTableGroup, global.GormDBGroup)
	zap.L().Info("数据库表适配完成!")
	// 初始化邮件服务器链接
	initialize.InitMail(global.Config.MailConfig)
	zap.L().Info("邮件客户端初始化成功")
	// 设置工作目录
	global.Config.ServerConfig.WorkDir = global.WorkDir
	// 初始化Http服务器
	s.server = initialize.InitHttpServer(global.Config.ServerConfig, global.Redis)
	return nil
}

// runHttpServer
// @Date 2023-01-11 17:47:36
// @Description: 运行Http服务器
func (s *Server) runHttpServer() {
	defer s.ShutDown()
	zap.L().Info("Http服务器正常运行",
		zap.String("端口", global.Config.ServerConfig.Addr()),
		zap.String("工作目录", global.Config.ServerConfig.WorkDir))
	err := s.server.ListenAndServe()
	if err != nil {
		zap.L().Error("应用启动失败，即将关闭", zap.Error(err))
	} else {
	}
}

// Run
// @Date: 2023-01-08 22:17:38
// @Description: 应用启动
// @Receiver: s *Server
func (s *Server) run() {
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
	initialize.CloseGormGroup(global.GormDBGroup)
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
	s.pareConfigPath()
	if fileutils.IsExist(s.ConfigPath) {
		s.run()
	} else {
		s.generateDefaultConfig()
	}
}

// generateDefaultConfig
// @Date 2023-01-13 19:14:38
// @Description: 创建默认配置文件
func (s *Server) generateDefaultConfig() {
	// 读取默认的配置文件内容
	templateCfg, err := resource.ResourceFS.ReadFile(defaultConfig)
	if err != nil {
		log.Panicln("默认配置文件读取失败:"+defaultConfig, err)
	}
	// 将默认内容写入配置文件路径内
	if err := os.WriteFile(s.ConfigPath, templateCfg, os.FileMode(0644)); err != nil {
		log.Panicln("默认配置创建异常:"+defaultTargetConfigPath, err)
	}
	log.Printf("默认配置文件已创建在%s，请完善配置文件内容后再次启动应用程序", s.ConfigPath)
}

// pareConfigPath
// @Date 2023-01-13 20:11:47
// @Return string
// @Description: 解析命令行中的配置文件参数
func (s *Server) pareConfigPath() {
	var configPath string
	flag.StringVar(&configPath, "cfg", defaultTargetConfigPath, "服务器配置文件")
	flag.Parse()
	// 如果配置文件格式不正确
	if !strings.Contains(configPath, defaultTargetConfigPath) {
		fmt.Printf("路径: %s 不是正确的配置文件路径，请输入正确的配置文件路径", configPath)
		os.Exit(0)
	}

	if !filepath.IsAbs(configPath) {
		s.ConfigPath = filepath.Join(global.WorkDir, configPath)
	} else {
		s.ConfigPath = configPath
	}
}
