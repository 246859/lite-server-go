package config

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	Mysql              = "mysql"
	PostgreSql         = "postgreSql"
	MysqlLoadConfigErr = errors.New("mysql load config failed")
	RedisLoadConfigErr = errors.New("redis load config failed")
)

type GormDBGroup = map[string]*gorm.DB
type DataBaseConfigGroup = map[string]*DataBaseConfig

// DataBaseConfig
// @Date: 2023-01-08 19:10:05
// @Description: Mysql配置
type DataBaseConfig struct {
	Enable      bool   `yaml:"enable"`
	Type        string `yaml:"type"`
	Ip          string `yaml:"ip"`
	Port        string `yaml:"port"`
	DbName      string `yaml:"dbName"`
	Config      string `yaml:"config"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	MaxIldeCons int    `yaml:"maxIldeCons"`
	MaxOpenCons int    `yaml:"maxOpenCons"`
}

type MysqlConfig struct {
	C *DataBaseConfig
}

func (m MysqlConfig) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", m.C.Username, m.C.Password, m.C.Ip, m.C.Port, m.C.DbName, m.C.Config)
}

// RedisConfig
// @Date: 2023-01-08 19:10:33
// @Description: Redis配置
type RedisConfig struct {
	Options  *redis.Options
	Ip       string
	Port     string
	Password string
}

func (r RedisConfig) RedisOptions() *redis.Options {
	r.Options.Addr = r.Ip + ":" + r.Port
	r.Options.Password = r.Password
	return r.Options
}
