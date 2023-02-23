package interact

import "gorm.io/gorm"

// Like
// @Date 2023-02-23 20:25:46
// @Description: 点赞信息表
type Like struct {
	Like int `json:"like" gorm:"comment:点赞数量;"`
	gorm.Model
}
