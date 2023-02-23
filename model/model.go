package model

import (
	"github.com/246859/lite-server-go/model/article"
	"github.com/246859/lite-server-go/model/interact"
	"github.com/246859/lite-server-go/model/sys"
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
	SystemTableList = []interface{}{
		sys.SystemUser{},
		article.Article{},
		article.ArticleComment{},
		article.ArticleLike{},
		interact.Comment{},
		interact.Like{},
	}
)
