package file

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Name string `gorm:"comment:文件名称;"`
	Type string `gorm:"comment:文件类型;"`
	Dir  string `gorm:"comment:文件目录;"`
	Path string `gorm:"comment:文件路径;"`
	Url  string `gorm:"comment:文件映射路径;"`
}

type FileInfo struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Url  string `json:"url"`
}
