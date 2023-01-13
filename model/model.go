package model

import "liteserver/model/sys"

type TableGroup = map[string]TableList
type TableList = []interface{}

var (
	ModelTableGroup = &TableGroup{
		"main": SystemTableList,
	}
	SystemTableList = []interface{}{sys.SystemUser{}}
)
