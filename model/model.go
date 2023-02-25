package model

type TableMeta interface {
	TableComment() string
	TableName() string
}

type TableGroup = map[string]TableList
type TableList = []TableMeta

var (
	ModelTableGroup = &TableGroup{
		"main": SystemTableList,
	}
	// SystemTableList
	// @Date: 2023-02-06 22:33:32
	// 系统表
	SystemTableList = []TableMeta{
		SystemUser{},
		Class{},
		File{},

		Article{},
		ArticleComment{},

		Share{},
		ShareComment{},

		Comment{},
		Reply{},
		Like{},
	}
)
