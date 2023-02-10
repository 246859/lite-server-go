package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

// ResponseWriterWrapper
// @Date 2023-02-10 19:56:35
// @Description: gin自带的writer不能直接读取响应体
// 这里用组合简单包装一下，用一个Buffer来存响应体的内容
type ResponseWriterWrapper struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (r *ResponseWriterWrapper) Write(b []byte) (int, error) {
	r.Body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (r *ResponseWriterWrapper) WriteString(s string) (int, error) {
	r.Body.WriteString(s)
	return r.ResponseWriter.WriteString(s)
}

func (r *ResponseWriterWrapper) String() string {
	return r.Body.String()
}

func NewResponseWriterWrapper(ctx *gin.Context) *ResponseWriterWrapper {
	return &ResponseWriterWrapper{
		ResponseWriter: ctx.Writer,
		Body:           bytes.NewBufferString(""),
	}
}
