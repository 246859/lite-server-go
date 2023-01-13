package initialize

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"liteserver/config"
	"liteserver/i18n"
	"liteserver/utils/fileutils"
	"os"
	"path/filepath"
	"strings"
)

// InitI18nInfo
// @Date 2023-01-12 17:22:11
// @Description: 初始化应用国际化信息
func InitI18nInfo(cfg *config.I18nConfig) *i18n.I18nLocale {
	langdir := fileutils.JoinPath(cfg.Dir)
	// 创建文件夹
	fileutils.MustMkdir(langdir)
	dir, _ := os.ReadDir(langdir)
	locale := i18n.I18nLocale{}
	// 读取目录下语言文件
	for _, entry := range dir {
		entryPath := filepath.Join(langdir, entry.Name())
		if !entry.IsDir() && strings.Contains(entryPath, cfg.Suffix) {
			localeName := strings.ReplaceAll(entry.Name(), cfg.Suffix, "")
			// 解析render
			locale[localeName] = resolveRender(entryPath)
		}
	}
	return &locale
}

// resolveRender
// @Date 2023-01-12 17:26:56
// @Param path string
// @Return i18n.LocaleRender
// @Description: 根据文件路径解析一个 i18n.LocaleRender
func resolveRender(path string) i18n.LocaleRender {
	render := viper.New()
	render.SetConfigFile(path)
	err := render.ReadInConfig()
	if err != nil {
		zap.L().Panic("国际化语言文件加载异常",
			zap.String("文件路径", path),
			zap.Error(err))
	}
	return render
}
