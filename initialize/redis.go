package initialize

import (
	"context"

	"github.com/go-redis/redis/v8"
	"liteserver/config"
)

// InitRedis
// @Date: 2023-01-09 11:07:45
// @Description: 初始化Redis客户端连接
// @Param: config *config.RedisConfig
// @Return: *redis.Client
func InitRedis(config *config.RedisConfig) *redis.Client {
	client := redis.NewClient(config.RedisOptions())
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return client
}
