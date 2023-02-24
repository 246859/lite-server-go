package response

import "github.com/246859/lite-server-go/model"

// FileInfo
// @Date 2023-02-23 23:04:26
// @Description: 文件展示信息
type FileInfo struct {
	model.FileTN `json:"_"`
	Name         string `json:"name" label:"文件名称"`
	Type         string `json:"type" label:"文件类型"`
	Url          string `json:"url" label:"文件url"`
}
