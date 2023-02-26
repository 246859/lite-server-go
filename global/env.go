package global

import (
	"github.com/246859/lite-server-go/config"
	"github.com/246859/lite-server-go/i18n"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

var (
	Redis       *redis.Client
	GormDBGroup *config.GormDBGroup
	Viper       *viper.Viper
	Logger      *zap.Logger
	I18nLocale  *i18n.I18nLocale
	WorkDir     string
)

// DB
// @Date: 2023-01-09 10:41:36
// @Description: 获取默认的GormDB
// @Return: *gorm.DB
func DB() *gorm.DB {
	return DBName("main")
}

func DBName(name string) *gorm.DB {
	if db, ok := (*GormDBGroup)[name]; ok {
		if Config.ServerConfig.Mode == "debug" {
			return db.Debug()
		}
		return db
	}
	return nil
}

func Model(ele any) *gorm.DB {
	return DB().Model(ele)
}

// I18nRaw
// @Date 2023-02-06 21:58:50
// @Param key string
// @Param locale language.Tag
// @Return string
// @Method
// @Description: 获取国际化信息
func I18nRaw(key string, locale language.Tag) string {
	return I18nLocale.GetWithRaw(key, locale)
}

func I18nRawCN(key string) string {
	return I18nRaw(key, language.Chinese)
}

func I18nDataCN(key string, data ...any) string {
	return I18nData(key, language.Chinese, data...)
}

// I18nData
// @Date 2023-02-06 22:02:31
// @Param key string
// @Param locale language.Tag
// @Param data ...any
// @Return string
// @Method
// @Description: 获取带参数的国际化信息
func I18nData(key string, locale language.Tag, data ...any) string {
	return I18nLocale.GetWithData(key, locale, data...)
}
