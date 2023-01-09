package server

import (
	"errors"
	"flag"

	"github.com/fsnotify/fsnotify"
	"liteserver/config"
	"liteserver/global"
	"liteserver/initialize"
)

type Server struct {
	ConfigPath string
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
		return errors.New("应用环境加载异常")
	}
	global.Logger = initialize.InitZap(global.Config.ZapConfig)
	global.Redis = initialize.InitRedis(global.Config.RedisConfig)
	global.GormDBGroup = initialize.InitGorm(global.Config.DataBaseConfig)
	return nil
}

// Run
// @Date: 2023-01-08 22:17:38
// @Description: 应用启动
// @Receiver: s *Server
func (s *Server) Run() {
	// 加载配置文件
	err := s.loadConfig()
	if err != nil {
		panic(err)
	}

	// 加载应用环境
	err = s.loadEnv()
	if err != nil {
		panic(err)
	}
	s.Refresh()
}

func (s *Server) Refresh() {
	global.Viper.WatchConfig()
	global.Viper.OnConfigChange(func(in fsnotify.Event) {
		// 加载应用环境
		err := s.loadEnv()
		if err != nil {
			panic(err)
		}
	})
}

// RunWithFlag
// @Date: 2023-01-08 22:22:37
// @Description: 使用命令行标志输入配置文件路径
// @Receiver: s *Server
func (s *Server) RunWithFlag() {
	configPath := flag.String("cfg", "./config.yml", "服务器配置文件")
	flag.Parse()
	s.ConfigPath = *configPath
	s.Run()
}

// RunWithDefault
// @Date: 2023-01-08 22:22:28
// @Description: 使用默认的配置文件路径
// @Receiver: s *Server
func (s *Server) RunWithDefault() {
	s.ConfigPath = "./config.yml"
	s.Run()
}
