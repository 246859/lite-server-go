package test

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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

func TestDefaultConfig(t *testing.T) {
	temp, err := os.Open("../template/defaultConfig.yml")
	fmt.Println(err)
	dist, err := os.Create("./config2.yml")
	fmt.Println(err)
	fmt.Println(io.Copy(dist, temp))
}
