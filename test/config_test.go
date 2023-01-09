package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Human interface {
	Talk() string
}

type Man struct {
	Name string `json:"name"`
}

func (m Man) Talk() string {
	return "man"
}

func TestLoadConfig(t *testing.T) {
	str := `{name:"jack"}`
	var man Human
	man = &Man{}
	json.Unmarshal([]byte(str), &man)
	fmt.Println(man.Talk())
}
