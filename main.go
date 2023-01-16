package main

import (
	"liteserver/server"
)

// main
// @Date: 2023-01-08 19:43:26
// @Description: 服务器启动函数
func main() {
	liteServer := new(server.Server)
	liteServer.RunWithFlag()
}
