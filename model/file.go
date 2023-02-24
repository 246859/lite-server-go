package model

import "gorm.io/gorm"

type File struct {
	FileTN
	User   SystemUser `gorm:"foreignKey:UserId;constraint:onUpdate:RESTRICT,onDelete:CASCADE"`
	UserId uint       `json:"userId" gorm:"comment:用户ID;"`
	Name   string     `gorm:"comment:文件名称;"`
	Type   string     `gorm:"comment:文件类型;"`
	Dir    string     `gorm:"comment:文件目录;"`
	Path   string     `gorm:"comment:文件路径;"`
	Url    string     `gorm:"comment:文件映射路径;"`
	gorm.Model
}

type FileTN struct{}

func (FileTN) TableName() string {
	return "files"
}
