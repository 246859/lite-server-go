package test

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestV5(t *testing.T) {
	size := int(1e6)
	mp := make(map[string]int, size)
	for i := 0; i < size; i++ {
		str := uuid.NewV1().String()
		if _, ok := mp[str]; ok {
			mp[str]++
			fmt.Println("出现重复")
			fmt.Println(str, mp[str])
		} else {
			mp[str]++
		}
	}
	fmt.Println(size == len(mp))
}
