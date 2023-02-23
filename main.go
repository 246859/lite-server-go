package main

import "github.com/246859/lite-server-go/server"

// main
// @Date: 2023-01-08 19:43:26
// @Description: 服务器启动函数
func main() {
	liteServer := new(server.Server)
	liteServer.RunWithFlag()
}
