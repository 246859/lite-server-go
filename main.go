package main

import (
	"liteserver/server"
	"liteserver/utils"
)

// main
// @Date: 2023-01-08 19:43:26
// @Description: 服务器启动函数
func main() {
	utils.LogBanner()
	liteServer := server.Server{}
	liteServer.RunWithFlag()
}
