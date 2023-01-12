package i18n

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
)

var (
	MsgNotExist    = errors.New("message not exit")
	LocaleNotExist = errors.New("locale not exist")
)

type LocaleRender = *viper.Viper

// I18nLocale
// @Date: 2023-01-12 15:17:58
// 国际信息
type I18nLocale map[string]LocaleRender

// GetWithData
// @Date 2023-01-12 16:37:57
// @Param key string
// @Param data any
// @Param locale language.Tag
// @Description: 根据key和locale和data来渲染语言信息
func (l I18nLocale) GetWithData(key string, locale language.Tag, data ...any) (string, error) {
	render, err := l.getLocaleRender(locale)
	if err != nil {
		return "", err
	}
	res := l.withData(key, render, data...)
	return res, nil
}

// GetWithRaw
// @Date 2023-01-12 17:08:54
// @Param key string
// @Param locale language.Tag
// @Return string
// @Description: 不带任何参数直接获取语言信息
func (l I18nLocale) GetWithRaw(key string, locale language.Tag) string {
	render, err := l.getLocaleRender(locale)
	if err != nil {
		return ""
	}
	return l.withRaw(key, render)
}

// withData
// @Date 2023-01-12 17:11:40
// @Param key string
// @Param data any
// @Param render LocaleRender
// @Return string
// @Return error
// @Description: 根据render获取语言信息
func (l I18nLocale) withData(key string, render LocaleRender, data ...any) string {
	msg := render.GetString(key)
	return fmt.Sprintf(msg, data...)
}

// withRaw
// @Date 2023-01-12 17:10:51
// @Param key string
// @Param render LocaleRender
// @Return string
// @Description: 仅通过key值来获取语言信息
func (l I18nLocale) withRaw(key string, render LocaleRender) string {
	return render.GetString(key)
}

// getLocaleRender
// @Date 2023-01-12 16:53:52
// @Param locale language.Tag
// @Return LocaleRender
// @Return error
// @Description: 获取对应语言的render
func (l I18nLocale) getLocaleRender(locale language.Tag) (LocaleRender, error) {
	v, ok := l[locale.String()]
	if !ok {
		return nil, LocaleNotExist
	}
	return v, nil
}
