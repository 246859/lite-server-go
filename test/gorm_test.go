package test

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

var db *gorm.DB

func init() {
	open, err := gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?%s",
		"root", "wyh246859", "192.168.48.134", "3306", "liteserver", "charset=utf8mb4&parseTime=True&loc=Local")))
	if err != nil {
		panic(err)
	} else {
		db = open
	}
}

type ArticleComment struct {
	ID      uint `gorm:"primaryKey;"`
	UserId  uint
	User    User `gorm:"foreignKey:UserId"`
	Content string
	Replies []Reply `gorm:"foreignKey:CommentId;"`
}

type User struct {
	UserName
	ID       uint `gorm:"primaryKey;"`
	Username string
}

type UserName struct {
}

func (UserName) TableName() string {
	return "users"
}

type Reply struct {
	User      User `gorm:"foreignKey:UserId"`
	UserId    uint
	ID        uint `gorm:"primaryKey;"`
	CommentId uint
	Content   string
}

func TestArticleCommentReplyList(t *testing.T) {
	var data []ArticleComment
	db.Table("comments").Preload("User").Preload("Replies").Preload("Replies.User").Find(&data)
	fmt.Printf("%+v\n", data)
}
