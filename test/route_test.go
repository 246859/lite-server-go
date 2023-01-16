package test

import (
	"fmt"
	"liteserver/router"
	"testing"
)

func TestRouter(t *testing.T) {
	group := router.LiteRouter.InitGroup()
	fmt.Println(group)
}
