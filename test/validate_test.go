package test

import (
	"fmt"
	"github.com/246859/lite-server-go/middleware"

	"testing"
)

type Person struct {
	Username string `binding:"required" label:"用户名"`
	Age      string `binding:"required" label:"年龄"`
}

func TestTrans(t *testing.T) {
	translator := middleware.UniverseTranslator{}
	var persons []Person
	persons = append(persons, Person{
		Username: "",
		Age:      "",
	})
	// 用户名为必填字段, 年龄为必填字段,
	fmt.Println(translator.ValidateStruct(persons))
}
