package middleware

import (
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
	"sync"
)

func UniverseValidateTranslator() *UniverseTranslator {
	return &UniverseTranslator{}
}

// UniverseTranslator
// @Date 2023-02-09 14:14:12
// @Method
// @Description: Gin自带的Validator错误消息不支持国际化
// 于是需要自定义validator来替换
type UniverseTranslator struct {
	onece     sync.Once
	validator *validator.Validate
	translate *ut.Translator
}

func (u *UniverseTranslator) Engine() any {
	u.lazyInit()
	return u.validator
}

// lazyInit
// @Date 2023-02-09 15:33:26
// @Method
// @Description: 懒加载
func (u *UniverseTranslator) lazyInit() {
	u.onece.Do(func() {
		u.validator = validator.New()
		u.validator.SetTagName("binding")
		if err := u.translator(); err != nil {
			u.validator = nil
			return
		}
		u.validator.RegisterTagNameFunc(customTagNameFunc)
	})
}

// ValidateStruct
// @Date 2023-02-09 15:37:39
// @Param obj any
// @Return error
// @Method
// @Description: 验证结构体，只接收结构体类型，指向结构体的指针，或者结构体切片/数组
func (u *UniverseTranslator) ValidateStruct(obj any) error {
	if obj == nil {
		return nil
	}
	u.lazyInit()
	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Ptr:
		return u.ValidateStruct(value.Elem().Interface())
	case reflect.Struct:
		return u.validateStruct(obj)
	case reflect.Slice, reflect.Array:
		length := value.Len()
		validateErrs := make(binding.SliceValidationError, 0, length)
		for i := 0; i < length; i++ {
			if err := u.ValidateStruct(value.Index(i).Interface()); err != nil {
				validateErrs = append(validateErrs, err)
			}
		}
		if len(validateErrs) == 0 {
			return nil
		}
		return validateErrs
	default:
		return nil
	}
}

// validateStruct
// @Date 2023-02-09 15:39:23
// @Param obj any
// @Return error
// @Method
// @Description: 结构体字段验证
func (u *UniverseTranslator) validateStruct(obj any) error {
	// 错误类型断言
	if err := u.validator.Struct(obj); err != nil {
		return u.errorTranslate(err.(validator.ValidationErrors))
	}
	return nil
}

// errorTranslate
// @Date 2023-02-09 15:39:47
// @Param errs validator.ValidationErrors
// @Return error
// @Method
// @Description: 错误消息翻译
func (u *UniverseTranslator) errorTranslate(errs validator.ValidationErrors) error {
	builder := strings.Builder{}
	for _, err := range errs {
		builder.WriteString(err.Translate(*u.translate))
		builder.WriteString(", ")
	}
	return errors.New(builder.String())
}

// translator
// @Date 2023-02-09 15:39:57
// @Return error
// @Method
// @Description: 构建翻译器
func (u *UniverseTranslator) translator() error {
	zh := zh.New()
	uni := ut.New(zh, zh)
	trans, found := uni.GetTranslator(zh.Locale())
	if !found {
		return errors.New("该语言不存在")
	}
	err := zh_trans.RegisterDefaultTranslations(u.validator, trans)
	if err != nil {
		return err
	}
	u.translate = &trans
	return nil
}

// customTagNameFunc
// @Date 2023-02-09 15:16:38
// @Param field reflect.StructField
// @Return string
// @Description: 自定字段标签名
func customTagNameFunc(field reflect.StructField) string {
	label := field.Tag.Get("label")
	if len(label) == 0 {
		return field.Name
	}
	return label
}
