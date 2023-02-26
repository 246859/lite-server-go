package utils

import (
	"fmt"
)

// WrapSimpleError
// @Date 2023-02-26 21:00:34
// @Param msg string
// @Param err error
// @Return error
// @Description: 简单的包装错误
func WrapSimpleError(msg string, err error) error {
	return fmt.Errorf("%s:%w", msg, err)
}
