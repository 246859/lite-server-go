package model

import (
	"liteserver/model/sys/sysrep"
)

type TableGroup = map[string]TableList
type TableList = []interface{}

var (
	ModelTableGroup = &TableGroup{
		"main": SystemTableList,
	}
	// SystemTableList
	// @Date: 2023-02-06 22:33:32
	// 系统表
	SystemTableList = []interface{}{sysrep.SystemUser{}}
)
