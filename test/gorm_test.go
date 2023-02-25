package test

import (
	"encoding/json"
	"fmt"
	"github.com/246859/lite-server-go/dao"
	"github.com/246859/lite-server-go/model"
	"github.com/246859/lite-server-go/model/request"
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
	list, err := dao.ArticleDao{}.GetArticleCommentList(db, request.PageInfo{
		Page: 1,
		Size: 1,
		Desc: 0,
	}, 1)
	fmt.Println(ToJsonString(list))
	fmt.Println(ToJsonString(err))
}

func ToJsonString(a any) string {
	marshal, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		return ""
	} else {
		return string(marshal)
	}
}

func TestStructField(t *testing.T) {
	var c model.Comment
	// gorm 没法根据结构体字段名来查询，得手动输入表字段名
	db.Model(model.CommentMeta{}).Where("? = ?", model.Comment{}.ID, 1).First(&c)
	fmt.Println(c)
}

func TestModel(t *testing.T) {
	// 这里犯了一个错误，方法名跟结构体字段同名
	fmt.Println(model.Reply{}.TableComment())
}
