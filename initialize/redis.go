package initialize

import (
	"github.com/go-redis/redis/v8"
	"liteserver/config"
)

// InitRedis
// @Date: 2023-01-09 11:07:45
// @Description: 初始化Redis客户端连接
// @Param: config *config.RedisConfig
// @Return: *redis.Client
func InitRedis(config *config.RedisConfig) *redis.Client {
	return redis.NewClient(config.RedisOptions())
}
