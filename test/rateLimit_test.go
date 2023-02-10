package test

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
	"time"
)

// TestRateLimit
// @Date 2023-02-10 20:58:31
// @Param t *testing.T
// @Method
// @Description: 令牌桶限流测试
func TestRateLimit(t *testing.T) {
	count := 0
	for {
		if count >= 100 {
			time.Sleep(time.Second)
			count = 0
		}
		bufferString := bytes.NewBufferString("")
		resp, _ := http.Get("http://localhost:8080/v1/ping")
		resp.Write(bufferString)
		fmt.Println(bufferString.String())
		count++
	}
}
