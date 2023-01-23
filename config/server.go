package config

import "time"

const (
	Tes = "test"
	Dev = "dev"
	Pro = "product"
)

// ServerConfig
// @Date 2023-01-11 16:54:02
// @Description: 应用配置
type ServerConfig struct {

	// Port
	// @Date 2023-01-11 16:54:44
	// @Description: 服务端口
	Port string `yaml:"port"`

	// Mode
	// @Date 2023-01-11 16:54:54
	// @Description: 运行模式 debug||test||release
	Mode string `yaml:"mode"`

	// ReadTimeout
	// @Date 2023-01-11 16:55:10
	// @Description: http请求读取超时时间
	ReadTimeout time.Duration `yaml:"readTimeout"`

	// WriteTimeout
	// @Date 2023-01-11 16:55:20
	// @Description: http写入响应的超时时间
	WriteTimeout time.Duration `yaml:"writeTimeout"`

	// IdleTimeout
	// @Date 2023-01-11 16:55:48
	// @Description: 连接复用的超时时间
	IdleTimeout time.Duration `yaml:"idleTimeout"`

	// ReadHeaderTimeout
	// @Date 2023-01-11 16:56:07
	// @Description: 请求头读取超时时间
	ReadHeaderTimeout time.Duration `yaml:"readHeaderTimeout"`

	// MultipartMemory
	// @Date 2023-01-11 16:56:24
	// @Description: Multipart表单的大小限制
	MultipartMemory int64 `yaml:"multipartMemory"`

	// MaxHeaderBytes
	// @Date 2023-01-11 16:57:04
	// @Description: 最大请求大小
	MaxHeaderBytes int `yaml:"maxHeaderBytes"`

	// WorkDir
	// @Date 2023-01-23 22:05:37
	// @Description: 服务器的工作目录
	WorkDir string

	// StaticDir
	// @Date 2023-01-23 22:08:32
	// @Description: 静态文件的存放路径
	StaticDir string `yaml:"static" mapstructure:"static"`
}

func (s *ServerConfig) Addr() string {
	return ":" + s.Port
}

func (s *ServerConfig) AddrHost(host string) string {
	return host + ":" + s.Port
}

// ReadTimeOut
// @Date 2023-01-11 16:18:53
// @Return time.Duration
// @Description: 读取Http请求的超时时间
func (s *ServerConfig) ReadTimeOut() time.Duration {
	return time.Duration(s.ReadTimeout * time.Second)
}

// WriteTimeOut
// @Date 2023-01-11 16:19:05
// @Return time.Duration
// @Description: 写入Http响应的超时时间
func (s *ServerConfig) WriteTimeOut() time.Duration {
	return time.Duration(s.WriteTimeout * time.Second)
}

// IdleTimeOut
// @Date 2023-01-11 16:19:19
// @Return time.Duration
// @Description: Keep alive Http连接复用的超时时间
func (s *ServerConfig) IdleTimeOut() time.Duration {
	return time.Duration(s.IdleTimeout * time.Second)
}

// ReadHeaderTimeOut
// @Date 2023-01-11 16:19:45
// @Return time.Duration
// @Description: 读取Http请求头的超时时间
func (s *ServerConfig) ReadHeaderTimeOut() time.Duration {
	return time.Duration(s.ReadHeaderTimeout * time.Second)
}
