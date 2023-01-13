package response

import (
	"github.com/gin-gonic/gin"
	"liteserver/global/code"
	"net/http"
)

// Response
// @Date 2023-01-12 21:36:46
// @Description: 统一全局响应体
type Response struct {
	// code
	// @Date 2023-01-12 21:51:51
	// @Description: 业务码
	Code code.BusinessCode `json:"code"`

	// Data
	// @Date 2023-01-12 21:52:08
	// @Description: 携带的数据
	Data interface{} `json:"data"`

	// Msg
	// @Date 2023-01-12 21:52:28
	// @Description: 基本信息
	Msg string `json:"msg"`
}

// NewResponse
// @Date 2023-01-12 22:25:03
// @Param c *gin.Context gin上下文
// @Param status int Http状态码
// @Param code code.BusinessCode 应用业务码
// @Param data interface{} 数据
// @Param msg string 信息
// @Description: 方便函数
func NewResponse(c *gin.Context, status int, code code.BusinessCode, data interface{}, msg string) {
	c.JSON(status, Response{
		code, data, msg,
	})
}

func Ok(c *gin.Context) {
	NewResponse(c, http.StatusOK, code.Success, nil, "")
}

func OkWithMsg(c *gin.Context, msg string) {
	NewResponse(c, http.StatusOK, code.Success, nil, msg)
}

func OkWithMsgAndData(c *gin.Context, data interface{}, msg string) {
	NewResponse(c, http.StatusOK, code.Success, data, msg)
}

func OkWithParams(c *gin.Context, code code.BusinessCode, data interface{}, msg string) {
	NewResponse(c, http.StatusOK, code, data, msg)
}

func Fail(c *gin.Context) {
	NewResponse(c, http.StatusBadRequest, code.BadOperation, nil, "")
}

func FailWithMsg(c *gin.Context, msg string) {
	NewResponse(c, http.StatusBadRequest, code.BadOperation, nil, msg)
}

func FailWithMsgAndData(c *gin.Context, data interface{}, msg string) {
	NewResponse(c, http.StatusBadRequest, code.BadOperation, data, msg)
}

func FailWithParams(c *gin.Context, code code.BusinessCode, data interface{}, msg string) {
	NewResponse(c, http.StatusBadRequest, code, data, msg)
}