package middleware

import (
	"errors"
	"liteserver/utils/response"
	"net"
	"net/http/httputil"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ZapLogger
// @Date: 2023-01-09 22:59:10
// @Description: 自定义的ZapLogger Gin中间件
// @Return: gin.HandlerFunc
func ZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始计时
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// 包装一下 gin.ResponseWriter
		responseWriterWrapper := NewResponseWriterWrapper(c)
		c.Writer = responseWriterWrapper
		// 处理
		c.Next()

		// 计算耗时
		cost := float64(time.Since(start).Nanoseconds()) / float64(time.Second)
		ip := c.ClientIP()
		userAgent := c.Request.UserAgent()
		method := c.Request.Method
		status := c.Writer.Status()
		ErrorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()
		bodySize := c.Writer.Size()
		responseContent := responseWriterWrapper.String()
		// 长度过大则不打印
		if len(responseContent) >= 300 {
			responseContent = "response length >= 300"
		}

		// 日志输出
		zap.L().Info("[Gin] HttpRequest",
			zap.String("Method", method),
			zap.Int("Status", status),
			zap.String("Path", path),
			zap.String("Query", query),
			zap.Float64("Cost", cost),
			zap.Int("BodySize", bodySize),
			zap.String("Ip", ip),
			zap.String("User-Agent", userAgent),
			zap.String("errors", ErrorMessage),
			zap.String("response-content", responseContent))
	}
}

// ZapCustomRecovery
// @Date 2023-01-11 15:23:16
// @Param logger *zap.Logger
// @Param handle gin.RecoveryFunc
// @Return gin.HandlerFunc
// @Description: 自定义Zap整合Gin Recovery中间件
func ZapCustomRecovery(logger *zap.Logger, handle gin.RecoveryFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Logger
		defer func() {
			if err := recover(); err != nil {
				// 检查断开的连接，因为其并不是导致panic堆栈跟踪的必要条件
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				// 提取请求信息
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				headers := strings.Split(string(httpRequest), "\r\n")
				for idx, header := range headers {
					current := strings.Split(header, ":")
					if current[0] == "Authorization" {
						headers[idx] = current[0] + ": *"
					}
				}
				headersToStr := strings.Join(headers, "\r\n")
				if brokenPipe {
					logger.Error("broken pipe",
						zap.String("header", headersToStr),
						zap.Any("panic", err))
				} else if gin.IsDebugging() {
					logger.Error("[Recovery] panic is recovered",
						zap.String("header", headersToStr),
						zap.Any("panic", err))
				} else {
					logger.Error("[Recovery] panic is recovered",
						zap.String("header", headersToStr),
						zap.Any("panic", err))
				}

				if brokenPipe {
					c.Abort()
					c.Error(err.(error))
				} else {
					handle(c, err)
				}
			}
		}()
		// 先处理请求
		c.Next()
	}
}

// ZapRecovery
// @Date 2023-01-11 15:27:36
// @Description: 默认的Zap Recovery
func ZapRecovery() gin.HandlerFunc {
	return ZapCustomRecovery(zap.L(), RecoveryHandler)
}

func RecoveryHandler(c *gin.Context, err any) {
	if err != nil {
		value := reflect.ValueOf(err)
		switch value.Kind() {
		case reflect.String:
			response.InternalErrorWithMsg(c, value.String())
		case reflect.Interface:
			if e, ok := err.(error); ok {
				response.InternalErrorWithMsg(c, e.Error())
			} else {
				response.InternalError(c)
			}
		default:
			response.InternalError(c)
		}
	}
}
