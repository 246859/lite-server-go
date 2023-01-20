package uuidtool

import (
	uuid "github.com/satori/go.uuid"
	"math/rand"
)

var (
	NameSpaceNormal = uuid.UUID{108, 105, 116, 101, 45, 115, 101, 114, 118, 101, 114, 45, 103, 111, 127, 0}
)

// NewUUIDv5
// @Date 2023-01-20 20:54:59
// @Param name string
// @Return string
// @Method
// @Description: 生成UUIDv5，https://datatracker.ietf.org/doc/html/rfc4122#section-4.3
func NewUUIDv5() string {
	return uuid.NewV5(randomNamespace(), uuid.NewV1().String()).String()
}

// randomNamespace
// @Date 2023-01-20 21:05:30
// @Return uuid.UUID
// @Method
// @Description: 生成一个随机的命名空间
func randomNamespace() uuid.UUID {
	randomSeed()
	var ns uuid.UUID
	for i := 0; i < len(ns); i++ {
		ns[i] = byte(rand.Intn(128))
	}
	return ns
}

// randomName
// @Date 2023-01-20 21:11:57
// @Return string
// @Method
// @Description: 生成随机的名称
func randomName() string {
	randomSeed()
	length := rand.Intn(60) + 68
	var name []byte
	for i := 0; i < length; i++ {
		name = append(name, byte(rand.Intn(128)))
	}
	return string(name)
}

func randomSeed() {
	rand.Seed(rand.Int63())
}
