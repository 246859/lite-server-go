package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"liteserver/initialize"
)

var (
	Redis       *redis.Client
	GormDBGroup *initialize.GormDBGroup
	Viper       *viper.Viper
	Logger      *zap.Logger
)

// GetDefaultGormDB
// @Date: 2023-01-09 10:41:36
// @Description: 获取默认的GormDB
// @Return: *gorm.DB
func GetDefaultGormDB() *gorm.DB {
	if db, ok := (*GormDBGroup)["main"]; ok {
		return db
	}
	return nil
}
